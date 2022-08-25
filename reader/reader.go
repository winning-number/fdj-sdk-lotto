// Package reader read zip archive and csv file with a grace reader closer.
// It could record the downloaded files inside a folder if you want.
package reader

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Reader read from a zip archive to retreive a csv reader
//
//go:generate mockery --name=Reader --output=mocks --filename=reader.go --outpkg=mocks
type Reader interface {
	Close() error
	CSVReader() io.ReadCloser
}

type reader struct {
	zipReader  *zip.Reader
	zipCloser  *zip.ReadCloser
	csvContent io.ReadCloser
}

type Option struct {
	OutputZipFile string
	OutputCSVFile string
}

func New(zipReader io.ReadCloser, option Option, fileName string) (Reader, error) {
	var err error

	driver := &reader{}
	if err := driver.zipSource(zipReader, option, fileName); err != nil {
		return nil, err
	}
	if err = driver.csvSource(option, fileName); err != nil {
		return nil, err
	}

	return driver, nil
}

func (r *reader) Close() error {
	var errMessage string

	if r.zipCloser != nil {
		if errZipCloser := r.zipCloser.Close(); errZipCloser != nil {
			errMessage += errZipCloser.Error()
		}
	}
	if r.csvContent != nil {
		if errCSVContent := r.csvContent.Close(); errCSVContent != nil {
			errMessage += errCSVContent.Error()
		}
	}
	if errMessage != "" {
		return errors.New(errMessage)
	}

	return nil
}

func (r reader) CSVReader() io.ReadCloser {
	return r.csvContent
}

func (r *reader) zipSource(zipReader io.ReadCloser, option Option, fileName string) error {
	var err error

	if zipReader == nil {
		return ErrInvalidReaderInput
	}
	// Create a zipReader from the ReadCloser without save the body
	if option.OutputZipFile == "" {
		var body []byte

		if body, err = io.ReadAll(zipReader); err != nil {
			return err
		}
		if r.zipReader, err = zip.NewReader(bytes.NewReader(body), int64(len(body))); err != nil {
			return err
		}

		return nil
	}

	// Write the compressed body inside a zipFile
	zipPath := fmt.Sprintf("%s/%s", option.OutputZipFile, fileName)
	if err = r.copyZipSource(zipReader, zipPath); err != nil {
		return err
	}

	// Create a new zip.Reader
	if r.zipCloser, err = zip.OpenReader(zipPath); err != nil {
		return err
	}
	r.zipReader = &r.zipCloser.Reader

	return nil
}

func (r *reader) copyZipSource(zipReader io.ReadCloser, zipPath string) error {
	var err error
	var zipFile *os.File

	if zipFile, err = os.Create(zipPath); err != nil {
		return err
	}
	defer zipFile.Close()
	if _, err = io.Copy(zipFile, zipReader); err != nil {
		return err
	}

	return nil
}

func (r *reader) csvSource(option Option, fileName string) error {
	var err error
	var zipFile *zip.File

	nbFile := len(r.zipReader.File)
	if nbFile != 1 {
		errMessage := fmt.Sprintf("%s expected number files %d instead of one",
			fileName, nbFile)

		return errors.Wrap(ErrInvalidArchive, errMessage)
	}

	// Create a copy of the csv file
	if option.OutputCSVFile != "" {
		if err = r.createCSVFile(option.OutputCSVFile, fileName); err != nil {
			return err
		}
	}

	zipFile = r.zipReader.File[0]
	if r.csvContent, err = zipFile.Open(); err != nil {
		return err
	}

	return nil
}

func (r *reader) createCSVFile(path string, fileName string) error {
	var err error
	var file *os.File
	var zipFile *zip.File
	var contentZip io.ReadCloser

	csvPath := fmt.Sprintf("%s/%s.csv", path, strings.Trim(fileName, ".zip"))
	if file, err = os.Create(csvPath); err != nil {
		return err
	}
	defer file.Close()

	zipFile = r.zipReader.File[0]
	buf := make([]byte, zipFile.FileInfo().Size())
	if contentZip, err = zipFile.Open(); err != nil {
		return err
	}
	defer contentZip.Close()
	if _, err := contentZip.Read(buf); err != nil {
		if !errors.Is(err, io.EOF) {
			return err
		}
	}
	if _, err = file.Write(buf); err != nil {
		return err
	}

	return nil
}
