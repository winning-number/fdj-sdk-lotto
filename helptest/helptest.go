// Package helptest provide some functions to help the unit tests
package helptest

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RemoveFile for removing ressources
func RemoveFile(t *testing.T, filePath string) {
	// remove created files, assert an error if file not found because should be created
	err := os.Remove(filePath)
	assert.NoError(t, err)
}

// OpenFile is a function helper to factorize the unit tests
// Just open a file and catch a test error if occurred
func OpenFile(t *testing.T, filepath string) *os.File {
	var file *os.File
	var err error

	if file, err = os.Open(filepath); err != nil {
		t.Error(err)
	}

	return file
}

type ZipContent struct {
	Name    string
	Content []byte
}

// CreateZipReader return the zip archive with any files inside a io.ReadClose
// the reader is typeof io.NopCloser and do not need to call the Close method
func CreateZipReader(t *testing.T, files ...ZipContent) io.ReadCloser {
	var err error
	var zipWriter *zip.Writer
	var w io.Writer
	var buf bytes.Buffer

	zipWriter = zip.NewWriter(&buf)
	defer zipWriter.Close()
	for i := range files {
		if w, err = zipWriter.Create(files[i].Name); err != nil {
			t.Error(err)
		}
		if _, err = w.Write(files[i].Content); err != nil {
			t.Error(err)
		}
	}

	if err = zipWriter.Flush(); err != nil {
		t.Error(err)
	}

	return io.NopCloser(&buf)
}

// CreateFakeHTTPClient to test the read zip file from a response request
// url use to validate the request path use by the driver.
// too, the url will be use to provide the namefile.
// body use to filled the http.Response.
func CreateFakeHTTPClient(t *testing.T, body io.ReadCloser, url string) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			assert.Equal(t, req.URL.String(), url)
			header := make(http.Header)
			header.Set("Content-Type", "application/zip")
			header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(url)))

			return &http.Response{
				StatusCode:   http.StatusOK,
				Status:       "200 OK",
				Body:         body,
				Header:       header,
				Uncompressed: false,
				Request:      nil,
			}
		}),
	}
}

// RoundTripFunc type to configure a response without execute the call http
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip method to implement http.RoundTripper interface
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
