package helptest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// file and folder use for the unit tests
const (
	testFolder = "../.ignore"
	testFile   = "test-file"
)

func TestRemoveFile(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		var file *os.File
		var err error

		path := fmt.Sprintf("%s/%s", testFolder, testFile)
		if file, err = os.Create(path); err != nil {
			t.Error(err)
		}
		defer file.Close()

		// no assert because assertion(s) are done inside the RemoveFile function
		RemoveFile(t, path)
	})
}

func TestOpenFile(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		var file *os.File
		var err error

		path := fmt.Sprintf("%s/%s", testFolder, testFile)
		if file, err = os.Create(path); err != nil {
			t.Error(err)
		}
		if err = file.Close(); err != nil {
			t.Error(err)
		}

		file = OpenFile(t, path)
		if assert.NotNil(t, file) {
			assert.Equal(t, path, file.Name())
			assert.NoError(t, file.Close())
		}
	})
}

func TestCreateZipReader(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		input := ZipContent{
			Content: []byte(`content zipped`),
			Name:    `filename`,
		}
		reader := CreateZipReader(t, input)

		if assert.NotNil(t, reader) {
			assert.NoError(t, reader.Close())
		}
	})
}

func TestCreateFakeHTTPClient(t *testing.T) {
	t.Run("Should be ok", func(t *testing.T) {
		url := "/path"

		expectedBody := []byte(`it's a unit test`)
		expectedHeader := make(http.Header)
		expectedHeader.Set("Content-Type", "application/zip")
		expectedHeader.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(url)))

		body := io.NopCloser(bytes.NewReader(expectedBody))
		httpClient := CreateFakeHTTPClient(t, body, url)
		request, err := httpClient.Get(url) //nolint:noctx // not need to use ctx inside unit test
		if assert.NoError(t, err) {
			assert.EqualValues(t, http.StatusOK, request.StatusCode)
			assert.EqualValues(t, expectedHeader, request.Header)
			buf := make([]byte, len(expectedBody))
			var err error
			if _, err = request.Body.Read(buf); err != nil {
				t.Error(err)
			}
			assert.EqualValues(t, expectedBody, buf)
			assert.NoError(t, request.Body.Close())
		}
	})
}
