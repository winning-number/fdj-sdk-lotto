// Package lotto is package to get all lotto draws from the fdj history
// see https://www.fdj.fr/jeux-de-tirage/loto/historique
// Like an api, could get any archives to parse them into a common Draw object
// Filters could be get the draw by Day of draw, type of lotto ...
package lotto

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/winning-number/fdjapi-lotto/csvparser"
	"github.com/winning-number/fdjapi-lotto/httpclient"
	"github.com/winning-number/fdjapi-lotto/reader"
)

// sources zip files to get history draw(s)
const (
	GrandLoto       = "grandloto_201912.zip"
	GrandLotoNoel   = "lotonoel_201703.zip"
	SuperLoto199605 = "superloto_199605.zip"
	SuperLoto200810 = "superloto_200810.zip"
	SuperLoto201703 = "superloto_201703.zip"
	SuperLoto201907 = "superloto_201907.zip"
	Loto197605      = "loto_197605.zip"
	Loto200810      = "loto_200810.zip"
	Loto201703      = "loto_201703.zip"
	Loto201902      = "loto_201902.zip"
	Loto201911      = "loto_201911.zip" // last update
)

const (
	BasePath = "https://media.fdj.fr/static/csv/loto"
)

// DrawTypeKey for the csvparser.ContextRecord
const (
	DrawTypeKey = "DrawType"
)

// LoadAPISourceDisable allow to disable the download of specific source
// By default all sources should be download
type LoadAPISourceDisable struct {
	GrandLoto       bool
	GrandLotoNoel   bool
	SuperLoto199605 bool
	SuperLoto200810 bool
	SuperLoto201703 bool
	SuperLoto201907 bool
	Loto197605      bool
	Loto200810      bool
	Loto201703      bool
	Loto201902      bool
	Loto201911      bool
}

// LoadAPIOption allow to disable specific source from the load csv files (see: LoadAPISourceDisable)
// If OutputZipFile is define, a copy of the zip content should be write in the path
// If OutputCSVFile is define, a copy of the csv content should be wrtie in the path
type LoadAPIOption struct {
	SourceDisable LoadAPISourceDisable
	SaveSources   reader.Option
}

// Lotto interface could get the full history from the fdj archive or parse them directly from a csv
// LoadAPI and LoadCSV add draws inside the internal draws list and order them by date and id
//
//go:generate mockery --name=Lotto --output=mocks --filename=lotto.go --outpkg=mocks
type Lotto interface {
	// Read history from the FDJ api with filter to select the sources
	LoadAPI(option LoadAPIOption) error
	LoadCSV(r io.Reader, drawType DrawType, drawVersion DrawVersion) (csvparser.Warning, error)

	DrawCount(filter Filter) int
	Draws(filter Filter) []Draw
}

type lotto struct {
	draws      []Draw
	httpClient *http.Client
}

// New Lotto instance collect the Draw history
func New() (Lotto, error) {
	var driver lotto

	driver.httpClient = httpclient.New()

	return &driver, nil
}

func (l *lotto) LoadAPI(option LoadAPIOption) error {
	var err error

	if err = l.loadingSources(option); err != nil {
		return err
	}

	return nil
}

func (l *lotto) LoadCSV(r io.Reader, drawType DrawType, drawVersion DrawVersion) (csvparser.Warning, error) {
	var err error
	var warn csvparser.Warning

	conf := l.csvParseConfiguration(drawType)
	switch drawVersion {
	case DrawV0:
		conf.CreateObject = func() any { return &DrawCSV0{} }
	case DrawV1:
		conf.CreateObject = func() any { return &DrawCSV1{} }
	case DrawV2:
		conf.CreateObject = func() any { return &DrawCSV2{} }
	case DrawV3:
		conf.CreateObject = func() any { return &DrawCSV3{} }
	case DrawV4:
		conf.CreateObject = func() any { return &DrawCSV4{} }
	default:
		return nil, ErrDrawVersion
	}

	if warn, err = csvparser.CSVParse(r, conf); err != nil {
		return nil, err
	}

	return warn, nil
}

func (l *lotto) DrawCount(filter Filter) int {
	draws := l.Draws(filter)

	return len(draws)
}

func (l lotto) Draws(f Filter) []Draw {
	matchesDraws := []Draw{}

	for i := range l.draws {
		if !f.MatchWithDraw(&l.draws[i]) {
			continue
		}
		matchesDraws = append(matchesDraws, l.draws[i])
	}

	return matchesDraws
}

//nolint:funlen,gocognit,gocyclo // prefer keep the loading sources together
func (l *lotto) loadingSources(option LoadAPIOption) error {
	// ...
	var err error

	if !option.SourceDisable.GrandLoto {
		if err = l.loadingSource(option, GrandLoto, DrawGrandLottoType, DrawV3); err != nil {
			return err
		}
	}
	if !option.SourceDisable.GrandLotoNoel {
		if err = l.loadingSource(option, GrandLotoNoel, DrawXmasLottoType, DrawV3); err != nil {
			return err
		}
	}
	if !option.SourceDisable.SuperLoto199605 {
		if err = l.loadingSource(option, SuperLoto199605, DrawSuperLottoType, DrawV0); err != nil {
			return err
		}
	}
	if !option.SourceDisable.SuperLoto200810 {
		if err = l.loadingSource(option, SuperLoto200810, DrawSuperLottoType, DrawV2); err != nil {
			return err
		}
	}
	if !option.SourceDisable.SuperLoto201703 {
		if err = l.loadingSource(option, SuperLoto201703, DrawSuperLottoType, DrawV3); err != nil {
			return err
		}
	}
	if !option.SourceDisable.SuperLoto201907 {
		if err = l.loadingSource(option, SuperLoto201907, DrawSuperLottoType, DrawV3); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto197605 {
		if err = l.loadingSource(option, Loto197605, DrawLottoType, DrawV1); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto200810 {
		if err = l.loadingSource(option, Loto200810, DrawLottoType, DrawV2); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto201703 {
		if err = l.loadingSource(option, Loto201703, DrawLottoType, DrawV3); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto201902 {
		if err = l.loadingSource(option, Loto201902, DrawLottoType, DrawV3); err != nil {
			return err
		}
	}
	if !option.SourceDisable.Loto201911 {
		if err = l.loadingSource(option, Loto201911, DrawLottoType, DrawV4); err != nil {
			return err
		}
	}

	return nil
}

func (l *lotto) loadingSource(option LoadAPIOption, filePath string, drawType DrawType, drawVersion DrawVersion) error {
	var reader reader.Reader
	var err error
	var warn csvparser.Warning

	if reader, err = l.sourceReader(option, filePath); err != nil {
		return err
	}
	defer reader.Close()
	if warn, err = l.LoadCSV(reader.CSVReader(), drawType, drawVersion); err != nil {
		return err
	}
	if len(warn) > 0 {
		return errors.Wrap(ErrInvalidFDJSource, fmt.Sprintf("header with values: %v", warn))
	}

	return nil
}

func (l *lotto) sourceReader(option LoadAPIOption, filePath string) (reader.Reader, error) {
	var resp *http.Response
	var req *http.Request
	var r reader.Reader
	var err error

	url := fmt.Sprintf("%s/%s", BasePath, filePath)
	if req, err = http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil); err != nil {
		return nil, err
	}
	if resp, err = l.httpClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if r, err = reader.New(resp.Body, option.SaveSources, filePath); err != nil {
		return nil, err
	}

	return r, nil
}

/*
** Private method use to configure the csv parser config
** Allow to enrich the draw list in a lotto instance
 */

func (l *lotto) drawRecorder(rec any, context csvparser.ContextRecord) error {
	var draw Draw
	var err error
	var drawType DrawType
	var obj DrawRecorder
	var ok bool
	var contextValue string

	if obj, ok = rec.(DrawRecorder); !ok {
		return ErrDrawTypeDecode
	}
	if contextValue, ok = context.Get(DrawTypeKey); !ok {
		return ErrDrawTypeKeyNotFound
	}
	if drawType, err = DrawTypeConverter(contextValue); err != nil {
		return err
	}
	if draw, err = obj.ConvertDraw(drawType); err != nil {
		return err
	}
	l.draws = append(l.draws, draw)

	return nil
}

func (l *lotto) csvParseConfiguration(drawType DrawType) csvparser.ParseConfig {
	conf := csvparser.ParseConfig{
		Context:        csvparser.NewContextRecord(),
		RecordObject:   l.drawRecorder,
		CommaSeparator: ';',
	}
	conf.Context.Set(DrawTypeKey, string(drawType))

	return conf
}
