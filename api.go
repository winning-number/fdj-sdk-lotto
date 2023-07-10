package lotto

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gofast-pkg/csv"
	xhttp "github.com/gofast-pkg/http"
	"github.com/gofast-pkg/zip"
	"github.com/winning-number/fdj-sdk-lotto/draw"
)

var (
	ErrInvalidDrawType       = errors.New("draw type instance is invalid")
	ErrInvalidCSVType        = errors.New("csv type instance is invalid")
	ErrInvalidResponseHTTP   = errors.New("response http is invalid")
	ErrInvalidContextDecoder = errors.New("context decoder is invalid")
	ErrInvalidHTTPRequest    = errors.New("http request is invalid")
	ErrNoContext             = errors.New("context is nil")
	ErrDrawTypeKeyNotFound   = errors.New("draw type key not found in the context recorder")
)

// API is the interface to load the history source.
// It can load the history source from the FDJ archive or from a file.
// To get a default api instance, use NewAPI function.
//
//go:generate mockery --name=API --output=mocks --filename=api.go --outpkg=mocks
type API interface {
	// LoadAPI loads the history source from the FDJ archive
	LoadSource(ctx context.Context, source []SourceInfo) error
	// DownloadSource downloads the history source from the FDJ archive
	DownloadSource(ctx context.Context, path string, source []SourceInfo) error
	// SourceUpdatedAt returns the last update time of the history source
	SourceUpdatedAt(ctx context.Context, source SourceInfo) (time.Time, error)
	// LoadFile loads the history source from a file
	// source is used to know the type of the source
	// File must be a csv file with ';' separator
	LoadFile(path string, source SourceInfo) error
	// NDraws returns the number of draws
	NDraws(filter Filter) int64
	// Draws returns the draws depending on the filter
	Draws(filter Filter, order draw.OrderType) []draw.Draw
}

type api struct {
	httpClient *http.Client
	draws      []draw.Draw
}

// NewAPI returns a new API
func NewAPI() API {
	return &api{
		httpClient: xhttp.NewClient(),
	}
}

func (a *api) LoadSource(ctx context.Context, source []SourceInfo) error {
	var err error
	var reader zip.Reader
	var buf []byte

	if ctx == nil {
		return ErrNoContext
	}
	for _, s := range source {
		if reader, err = a.requestSource(ctx, s); err != nil {
			return err
		}
		for i := 0; i < reader.NFiles(); i++ {
			if buf, err = reader.ContentFile(i); err != nil {
				return err
			}
			if err = a.parseCSV(bytes.NewBuffer(buf), s); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a api) DownloadSource(ctx context.Context, path string, source []SourceInfo) error {
	var err error
	var reader zip.Reader

	if ctx == nil {
		return ErrNoContext
	}
	for _, s := range source {
		if reader, err = a.requestSource(ctx, s); err != nil {
			return err
		}
		for i := 0; i < reader.NFiles(); i++ {
			if err = reader.WriteFile(i, path); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a api) SourceUpdatedAt(ctx context.Context, source SourceInfo) (time.Time, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	var t time.Time

	if ctx == nil {
		return time.Time{}, ErrNoContext
	}
	if req, err = http.NewRequestWithContext(ctx, http.MethodHead, source.URL(), nil); err != nil {
		return time.Time{}, errors.Join(err, ErrInvalidHTTPRequest)
	}
	if resp, err = a.httpClient.Do(req); err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return time.Time{}, errors.Join(
			ErrInvalidResponseHTTP,
			fmt.Errorf("with status %s", resp.Status))
	}
	if t, err = time.Parse(
		"Mon, 02 Jan 2006 15:04:05 MST",
		resp.Header.Get("Last-Modified"),
	); err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func (a *api) LoadFile(path string, source SourceInfo) error {
	var err error
	var file *os.File

	if file, err = os.Open(path); err != nil {
		return err
	}
	defer file.Close()

	if err = a.parseCSV(file, source); err != nil {
		return err
	}

	return nil
}

func (a api) NDraws(filter Filter) int64 {
	return int64(len(a.Draws(filter, draw.OrderNone)))
}

func (a api) Draws(filter Filter, order draw.OrderType) []draw.Draw {
	matchesDraws := []draw.Draw{}

	for i := range a.draws {
		if !filter.Match(&a.draws[i]) {
			continue
		}
		matchesDraws = append(matchesDraws, a.draws[i])
	}
	draw.OrderDraws(&matchesDraws, order)

	return matchesDraws
}

func (a *api) parseCSV(input io.Reader, source SourceInfo) error {
	var err error
	var csvReader csv.CSV
	var decoder csv.Decoder
	var warn csv.Warning

	if csvReader, err = csv.New(input, ';'); err != nil {
		return err
	}
	if decoder, err = csv.NewDecoder(csv.ConfigDecoder{
		NewInstanceFunc: func() any {
			return newInstanceFunc(source.Version)
		},
		SaveInstanceFunc: func(drawInstance any, decoder csv.Decoder) error {
			var saveInstErr error
			var d draw.Draw

			if d, saveInstErr = saveInstanceFunc(drawInstance, decoder); saveInstErr != nil {
				return err
			}
			if draw.DrawFinder(&a.draws, d) {
				return nil
			}
			a.draws = append(a.draws, d)

			return nil
		},
	}); err != nil {
		return err
	}
	decoder.ContextSet(keyDrawType, (string)(source.Type))

	if warn, err = csvReader.DecodeWithDecoder(decoder); err != nil {
		return err
	}
	if len(warn) > 0 {
		return errors.Join(ErrInvalidCSVType, fmt.Errorf("%v", warn))
	}

	return nil
}

func (a api) requestSource(ctx context.Context, source SourceInfo) (zip.Reader, error) {
	var err error
	var resp *http.Response
	var req *http.Request
	var reader zip.Reader

	if req, err = http.NewRequestWithContext(ctx, http.MethodGet, source.URL(), nil); err != nil {
		return nil, errors.Join(err, ErrInvalidHTTPRequest)
	}
	if resp, err = a.httpClient.Do(req); err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Join(
			ErrInvalidResponseHTTP,
			fmt.Errorf("with response status %s", resp.Status))
	}
	defer resp.Body.Close()

	if reader, err = zip.NewReader(resp.Body); err != nil {
		return nil, err
	}

	return reader, nil
}
