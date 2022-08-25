package lotto

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
** Function to help the unit tests
**
 */

// datatest files
const (
	superLottoTestFileV0     = "testdata/super-loto-v0.csv"
	superLottoTestFileV2     = "testdata/super-loto-v2.csv"
	superLottoTestFileV3     = "testdata/super-loto-v3.csv"
	grandLottoTestFileV3     = "testdata/grand-loto-v3.csv"
	grandLottoNoelTestFileV3 = "testdata/grand-loto-noel-v3.csv"
	classicLottoTestFileV1   = "testdata/classic-loto-v1.csv"
	classicLottoTestFileV2   = "testdata/classic-loto-v2.csv"
	classicLottoTestFileV3   = "testdata/classic-loto-v3.csv"
	classicLottoTestFileV4   = "testdata/classic-loto-v4.csv"
)

// helpOpenFile is a function helper to factorize the unit tests
// Just open a file and catch a test error if occurred
func helpOpenFile(t *testing.T, filepath string) *os.File {
	var file *os.File
	var err error

	if file, err = os.Open(filepath); err != nil {
		t.Error(err)
	}
	return file
}

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

// helpCreateLottoWithFakeHTTPClient
// expectedPath use to validate the request path use by the driver.
// too, the expectedPath will be use to provide the namefile.
// body use to filled the http.Response.
func helpCreateLottoWithFakeHTTPClient(t *testing.T, body io.ReadCloser, expectedPath string) *lotto {
	return &lotto{
		httpClient: &http.Client{
			Transport: helpRoundTripFunc(func(req *http.Request) *http.Response {
				assert.Equal(t, req.URL.String(), fmt.Sprintf("%s/%s", BasePath, expectedPath))
				header := make(http.Header)
				header.Set("Content-Type", "application/zip")
				header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", expectedPath))

				return &http.Response{
					StatusCode:   200,
					Status:       "200 OK",
					Body:         body,
					Header:       header,
					Uncompressed: false,
					Request:      nil,
				}
			}),
		},
	}
}

// helpRoundTripFunc type to configure a response without execute the call http
type helpRoundTripFunc func(req *http.Request) *http.Response

// RoundTrip method to implement http.RoundTripper interface
func (f helpRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// helpLoadAPIOptionSourceDisabled return LoadAPIOption without source enable
func helpLoadAPIOptionSourceDisabled() LoadAPIOption {
	return LoadAPIOption{
		SourceDisable: LoadAPISourceDisable{
			GrandLoto:       true,
			GrandLotoNoel:   true,
			SuperLoto199605: true,
			SuperLoto200810: true,
			SuperLoto201703: true,
			SuperLoto201907: true,
			Loto197605:      true,
			Loto200810:      true,
			Loto201703:      true,
			Loto201902:      true,
			Loto201911:      true,
		},
	}
}
