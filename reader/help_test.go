package reader

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
** Function to help the unit tests
**
 */

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

type zipContent struct {
	name    string
	content []byte
}

// helpCreateZipReader return the zip archive with any files inside a io.ReadClose
// the reader is typeof io.NopCloser and do not need to call the Close method
func helpCreateZipReader(t *testing.T, files ...zipContent) io.ReadCloser {
	var err error
	var zipWriter *zip.Writer
	var w io.Writer
	var buf bytes.Buffer

	zipWriter = zip.NewWriter(&buf)
	defer zipWriter.Close()
	for _, file := range files {
		if w, err = zipWriter.Create(file.name); err != nil {
			t.Error(err)
		}
		w.Write(file.content)
	}

	if err = zipWriter.Flush(); err != nil {
		t.Error(err)
	}

	return io.NopCloser(&buf)
}

// helpCloseSourceReader close the sources after unit tests
func helpCloseSourceReader(t *testing.T, driver Reader) {
	var source *reader
	var ok bool

	if source, ok = driver.(*reader); !ok {
		t.Error("SourceReader do not implemented by *sourceReader")
	}

	if source.csvContent != nil {
		source.csvContent.Close()
	}
	if source.zipCloser != nil {
		source.zipCloser.Close()
	}
}

// helpRemoveFile for removing ressources
func helpRemoveFile(t *testing.T, filePath string) {
	// remove created files, assert an error if file not found because should be created
	err := os.Remove(filePath)
	assert.NoError(t, err)
}
