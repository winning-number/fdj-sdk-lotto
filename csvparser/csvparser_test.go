package csvparser

import (
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

/*
** Test data stuff
 */

// testfiles
const (
	testFilePath      = "testdata/testfile.csv"
	testEmptyFilePath = "testdata/testfile_empty.csv"
	testWrongFilePath = "testdata/testfile_wrong_csv_format.csv"
)

type DataTestLong struct {
	Name      string `csv:"name"`
	Type      string `csv:"type"`
	MainColor string `csv:"main_color"`
	Size      string `csv:"size"`
}

func newDataTestLong() interface{} {
	return &DataTestLong{}
}

type DataTestShort struct {
	Name string `csv:"name"`
	Type string `csv:"type"`
}

func newDataTestShort() interface{} {
	return &DataTestShort{}
}

/*
** Unit tests
 */

func TestCSVParse(t *testing.T) {
	errObjConvertion := errors.New("bad type convertion for the recording")

	t.Run("Should return an error because config not setup", func(t *testing.T) {
		warn, err := CSVParse(nil, ParseConfig{})
		if assert.Error(t, err) {
			assert.Nil(t, warn)
			assert.EqualError(t, err, ErrBadConfiguration.Error())
		}
	})
	t.Run("Should return an error because reader is nil", func(t *testing.T) {
		warn, err := CSVParse(nil, ParseConfig{
			CreateObject:   newDataTestShort,
			RecordObject:   func(obj any, Context ContextRecord) error { return nil },
			Context:        NewContextRecord(),
			CommaSeparator: ';',
		})
		if assert.Error(t, err) {
			assert.Nil(t, warn)
			assert.EqualError(t, err, ErrNilReader.Error())
		}
	})
	t.Run("Should return an error because csvutils.NewDecoder fail to load an empty csv", func(t *testing.T) {
		var datas []DataTestLong
		file, err := os.Open(testEmptyFilePath)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		warn, err := CSVParse(file, ParseConfig{
			CreateObject:   newDataTestShort,
			RecordObject:   func(obj any, Context ContextRecord) error { return nil },
			Context:        NewContextRecord(),
			CommaSeparator: ';',
		})

		if assert.Error(t, err) {
			assert.Nil(t, warn)
			assert.ErrorIs(t, err, ErrNewDecoder)
			assert.Empty(t, datas)
		}
	})
	t.Run("Should return an error because decoder.Decode was fail", func(t *testing.T) {
		var datas []DataTestLong

		file, err := os.Open(testWrongFilePath)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		warn, err := CSVParse(file, ParseConfig{
			CreateObject:   newDataTestShort,
			RecordObject:   func(obj any, Context ContextRecord) error { return nil },
			Context:        NewContextRecord(),
			CommaSeparator: ';',
		})
		if assert.Error(t, err) {
			assert.Nil(t, warn)
			assert.ErrorIs(t, err, ErrOBJDecode)
			assert.Empty(t, datas)
		}
	})
	t.Run("Should return an error because the function recorder was fail", func(t *testing.T) {
		var data []DataTestLong

		file, err := os.Open(testFilePath)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		warn, err := CSVParse(file, ParseConfig{
			CreateObject:   newDataTestShort,
			RecordObject:   func(obj any, Context ContextRecord) error { return errors.New("one error") },
			Context:        NewContextRecord(),
			CommaSeparator: ';',
		})
		if assert.Error(t, err) {
			assert.Nil(t, warn)
			assert.ErrorIs(t, err, ErrRecorder)
			assert.Empty(t, data)
		}
	})
	t.Run("Should be ok without warning", func(t *testing.T) {
		var data []DataTestLong
		conf := ParseConfig{
			CreateObject: newDataTestLong,
			RecordObject: func(obj any, Context ContextRecord) error {
				v, ok := obj.(*DataTestLong)
				if !ok {
					t.Error(errObjConvertion)
				}
				data = append(data, *v)

				return nil
			},
			Context:        NewContextRecord(),
			CommaSeparator: ';',
		}

		expectedData := []DataTestLong{{
			Name:      "Root",
			Type:      "Dog",
			MainColor: "black",
			Size:      "big",
		}, {
			Name:      "Toto",
			Type:      "Human",
			MainColor: "blue",
			Size:      "small",
		}}

		file, err := os.Open(testFilePath)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		warn, err := CSVParse(file, conf)
		if assert.NoError(t, err) {
			assert.Empty(t, warn)
			assert.EqualValues(t, expectedData, data)
		}
	})
	t.Run("Should be ok with warning", func(t *testing.T) {
		var data []DataTestShort
		conf := ParseConfig{
			CreateObject: newDataTestShort,
			RecordObject: func(obj any, Context ContextRecord) error {
				v, ok := obj.(*DataTestShort)
				if !ok {
					t.Error(errObjConvertion)
				}
				data = append(data, *v)

				return nil
			},
			Context:        NewContextRecord(),
			CommaSeparator: ';',
		}

		expectedData := []DataTestShort{{
			Name: "Root",
			Type: "Dog",
		}, {
			Name: "Toto",
			Type: "Human",
		}}
		expectedWarn := Warning{
			"main_color": {"black", "blue"},
			"size":       {"big", "small"},
		}

		file, err := os.Open(testFilePath)
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		warn, err := CSVParse(file, conf)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedWarn, warn)
			assert.EqualValues(t, expectedData, data)
		}
	})
}
