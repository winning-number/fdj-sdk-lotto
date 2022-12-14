package reader

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/winning-number/fdj-sdk-lotto/helptest"
)

// files to test have a zip.ReadCloser
const (
	testReadCloserZipFile = "testdata/zipfile.zip"
)

// files and folder use for the tests file creation
const (
	testZipFile = "test-file.zip"
	testFolder  = "../.ignore"
	testCSVFile = "test-file.csv"
)

// helpCloseSourceReader close the sources after unit tests
func helpCloseSourceReader(t *testing.T, driver Reader) {
	var source *reader
	var ok bool
	var err error

	if source, ok = driver.(*reader); !ok {
		t.Error("SourceReader do not implemented by *sourceReader")
	}

	if source.csvContent != nil {
		if err = source.csvContent.Close(); err != nil {
			t.Error(err)
		}
	}
	if source.zipCloser != nil {
		if err = source.zipCloser.Close(); err != nil {
			t.Error(err)
		}
	}
}

func TestNew(t *testing.T) {

	t.Run("Should return an error because io.ReadCloser is nil", func(t *testing.T) {
		source, err := New(nil, Option{}, "no-name")

		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrInvalidReaderInput)
			assert.Nil(t, source)
		}
	})
	t.Run("Should return an error because io.ReadCloser is already closed", func(t *testing.T) {
		expectedErr := "file already closed"
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		if err := file.Close(); err != nil {
			t.Error(err)
		}
		source, err := New(file, Option{}, "no-name")

		if assert.Error(t, err) {
			assert.ErrorContains(t, err, expectedErr)
			assert.Nil(t, source)
		}
	})
	t.Run("Should return an error because reader is an invalid zip file", func(t *testing.T) {
		r := io.NopCloser(strings.NewReader(`invalid zip file`))
		source, err := New(r, Option{}, "no-name")

		if assert.Error(t, err) {
			assert.ErrorContains(t, err, "zip: not a valid zip file")
			assert.Nil(t, source)
		}
	})
	t.Run("Should return an error because folder path for the zip file do not exist", func(t *testing.T) {
		r := helptest.CreateZipReader(t, helptest.ZipContent{
			Content: []byte(`content zipped for unit testing`)})

		source, err := New(r, Option{
			OutputZipFile: ".ignore/folder-do-not-exist",
		}, "no-name")

		if assert.Error(t, err) {
			assert.ErrorContains(t, err, "no such file or directory")
			assert.Nil(t, source)
		}
	})
	t.Run("Should return an error because reader file contains any files", func(t *testing.T) {
		r := helptest.CreateZipReader(t,
			helptest.ZipContent{Name: "test"},
			helptest.ZipContent{Name: "test2"})

		source, err := New(r, Option{}, "no-name")

		if assert.Error(t, err) {
			assert.ErrorIs(t, err, ErrInvalidArchive)
			assert.Nil(t, source)
		}
	})
	t.Run("Should return an error because csvFile path do not exist for the creation", func(t *testing.T) {
		r := helptest.CreateZipReader(t, helptest.ZipContent{
			Content: []byte(`content zipped for unit testing`),
			Name:    "test"})

		source, err := New(r, Option{
			OutputCSVFile: ".ignore/folder-do-not-exist",
		}, "no-name")

		if assert.Error(t, err) {
			assert.ErrorContains(t, err, "no such file or directory")
			assert.Nil(t, source)
		}
	})
	t.Run("Should be ok with the creation csv file", func(t *testing.T) {
		defer helptest.RemoveFile(t, fmt.Sprintf("%s/%s", testFolder, testCSVFile))
		r := helptest.CreateZipReader(t, helptest.ZipContent{
			Content: []byte(`content zipped for unit testing`),
			Name:    "test"})

		source, err := New(r, Option{
			OutputCSVFile: testFolder,
		}, testZipFile)
		defer helpCloseSourceReader(t, source)

		if assert.NoError(t, err) {
			assert.NotNil(t, source)
		}
	})
	t.Run("Should be ok with the creation zip file", func(t *testing.T) {
		defer helptest.RemoveFile(t, fmt.Sprintf("%s/%s", testFolder, testZipFile))
		r := helptest.CreateZipReader(t, helptest.ZipContent{
			Content: []byte(`content zipped for unit testing`),
			Name:    "test"})

		source, err := New(r, Option{
			OutputZipFile: testFolder,
		}, testZipFile)
		defer helpCloseSourceReader(t, source)

		if assert.NoError(t, err) {
			assert.NotNil(t, source)
		}
	})
	t.Run("Should be ok without creation files", func(t *testing.T) {
		r := helptest.CreateZipReader(t, helptest.ZipContent{
			Content: []byte(`content zipped for unit testing`),
			Name:    "test"})

		source, err := New(r, Option{}, "no-name")
		defer helpCloseSourceReader(t, source)

		if assert.NoError(t, err) {
			assert.NotNil(t, source)
		}
	})
}

func TestReader_Close(t *testing.T) {
	t.Run("Should return an error because zipCloser are already closed", func(t *testing.T) {
		zipCloser, err := zip.OpenReader(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		if err := zipCloser.Close(); err != nil {
			t.Error(err)
		}
		r := &reader{
			zipCloser: zipCloser,
		}
		expectedErr := "file already closed"

		err = r.Close()
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, expectedErr)
		}
	})
	t.Run("Should return an error because csvReader are already closed", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		if err := file.Close(); err != nil {
			t.Error(err)
		}
		r := &reader{
			csvContent: file,
		}
		expectedErr := "file already closed"

		err = r.Close()
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, expectedErr)
		}
	})
	t.Run("Should be ok with no reader set", func(t *testing.T) {
		r := &reader{}
		err := r.Close()

		assert.NoError(t, err)
	})
	t.Run("Should be ok with ", func(t *testing.T) {
		zipCloser, err := zip.OpenReader(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		csvReader := io.NopCloser(strings.NewReader("coucou"))
		r := &reader{
			zipCloser:  zipCloser,
			csvContent: csvReader,
		}
		err = r.Close()

		assert.NoError(t, err)
	})
}

func TestReader_CSVReader(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		file, err := os.Open(testReadCloserZipFile)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		r := &reader{
			csvContent: file,
		}

		csvReader := r.CSVReader()
		assert.Equal(t, file, csvReader)
	})
}
