// Package csvparser is a package to parse a full reader inside a slice object.
// Allow of configure a recorder to transform directly csv object into worker object.
// Warn process allow to know if the struct is identical if the reader
package csvparser

import (
	"encoding/csv"
	"io"

	"github.com/jszwec/csvutil"
	"github.com/pkg/errors"
)

// CSFFileExtension for the csv file type
const (
	CSVFileExtension = ".csv"
)

// ParseConfig could to process a full csv reader with a recorder for each row
// CreateObject should return the concrete type to parse the CSV
// RecordObject function handler that is call after each decoded object
// CommaSeparator type to the header of the csv
// Context should be pass on any record handler to retreive data in CSVParse method
type ParseConfig struct {
	CreateObject   func() any
	RecordObject   func(obj any, Context ContextRecord) error
	CommaSeparator rune
	Context        ContextRecord
}

func (c ParseConfig) isValid() bool {
	if c.CreateObject == nil || c.RecordObject == nil || c.Context == nil {
		return false
	}

	return true
}

// CSVParse process the full reader parameter and record each entries inside
// the obj returned by the CreateObject of ParseConfig.
// Pass the context inside the ParseConfig in each RecordObject call
func CSVParse(r io.Reader, cfg ParseConfig) (Warning, error) {
	var err error
	var csvReader *csv.Reader
	var decoder *csvutil.Decoder
	var warn Warning

	if !cfg.isValid() {
		return nil, ErrBadConfiguration
	}
	if r == nil {
		return nil, ErrNilReader
	}
	csvReader = csv.NewReader(r)
	csvReader.Comma = cfg.CommaSeparator
	if decoder, err = csvutil.NewDecoder(csvReader); err != nil {
		return nil, errors.Wrap(ErrNewDecoder, err.Error())
	}
	decoder.DisallowMissingColumns = true

	if warn, err = csvDecode(decoder, cfg); err != nil {
		return nil, err
	}

	return warn, nil
}

func csvDecode(decoder *csvutil.Decoder, cfg ParseConfig) (Warning, error) {
	var err error
	warn := NewWarning()
	for {
		obj := cfg.CreateObject()
		if err := decoder.Decode(obj); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, errors.Wrap(ErrOBJDecode, err.Error())
		}
		// aggregate unused fields in a Warning
		fillWarning(decoder, warn)

		if err = cfg.RecordObject(obj, cfg.Context); err != nil {
			return nil, errors.Wrap(ErrRecorder, err.Error())
		}
	}

	return warn, nil
}

func fillWarning(decoder *csvutil.Decoder, warn Warning) {
	header := decoder.Header()
	for _, i := range decoder.Unused() {
		if header[i] == "" {
			continue
		}
		warn.addUnusedValues(header[i], decoder.Record()[i])
	}
}
