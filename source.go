package lotto

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
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

type SourceReader interface {
	Close() error
	CSVReader() io.ReadCloser
}

type sourceReader struct {
	zipReader  *zip.Reader
	zipCloser  *zip.ReadCloser
	csvContent io.ReadCloser
}

func (s *sourceReader) Close() error {
	var errMessage string

	if s.zipCloser != nil {
		if errZipCloser := s.zipCloser.Close(); errZipCloser != nil {
			errMessage += errZipCloser.Error()
		}
	}
	if s.csvContent != nil {
		if errCSVContent := s.csvContent.Close(); errCSVContent != nil {
			errMessage += errCSVContent.Error()
		}
	}
	if errMessage != "" {
		return errors.New(errMessage)
	}

	return nil
}

func (s sourceReader) CSVReader() io.ReadCloser {
	return s.csvContent
}

func newSourceReader(r io.ReadCloser, option LoadAPIOption, fileName string) (SourceReader, error) {
	var err error

	source := &sourceReader{}
	defer source.Close()
	if err := source.zipSource(r, option, fileName); err != nil {
		return nil, err
	}
	if err = source.csvSource(option, fileName); err != nil {
		return nil, err
	}

	return source, nil
}

func (s *sourceReader) zipSource(r io.ReadCloser, option LoadAPIOption, fileName string) error {
	var err error
	var zipFile *os.File

	// Create a zipReader from the ReadCloser without save the body
	if option.OutputZipFile == "" {
		var body []byte

		if body, err = io.ReadAll(r); err != nil {
			return err
		}
		if s.zipReader, err = zip.NewReader(bytes.NewReader(body), int64(len(body))); err != nil {
			return err
		}

		return nil
	}

	// Write the compressed body inside a zipFile
	zipPath := fmt.Sprintf("%s/%s", option.OutputZipFile, fileName)
	if err = func() error {
		if zipFile, err = os.Create(zipPath); err != nil {
			return err
		}
		defer zipFile.Close()
		if _, err = io.Copy(zipFile, r); err != nil {
			return err
		}

		return nil
	}(); err != nil {
		return err
	}

	// Create a new zip.Reader
	if s.zipCloser, err = zip.OpenReader(zipPath); err != nil {
		return err
	}
	s.zipReader = &s.zipCloser.Reader

	return nil
}

func (s *sourceReader) csvSource(option LoadAPIOption, fileName string) error {
	var err error
	var zipFile *zip.File

	nbFile := len(s.zipReader.File)
	if nbFile != 1 {
		errMessage := fmt.Sprintf("%s expected number files %d instead of one",
			fileName, nbFile)

		return errors.Wrap(ErrInvalidArchive, errMessage)
	}

	// Create a copy of the csv file
	if option.OutputCSVFile != "" {
		if err = s.createCSVFile(option.OutputCSVFile, fileName); err != nil {
			return err
		}
	}

	zipFile = s.zipReader.File[0]
	if s.csvContent, err = zipFile.Open(); err != nil {
		return err
	}

	return nil
}

func (s *sourceReader) createCSVFile(path string, fileName string) error {
	var err error
	var file *os.File
	var zipFile *zip.File
	var contentZip io.ReadCloser

	csvPath := fmt.Sprintf("%s/%s.csv", path, strings.Trim(fileName, ".zip"))
	if file, err = os.Create(csvPath); err != nil {
		return err
	}
	defer file.Close()

	zipFile = s.zipReader.File[0]
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
