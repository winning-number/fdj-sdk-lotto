package lotto

import "github.com/pkg/errors"

// error list
var (
	ErrDrawTypeKeyNotFound    = errors.New("draw type key not found in the context recorder")
	ErrDrawTypeDecode         = errors.New("wrong draw type to decode")
	ErrNoCreateObjectToDecode = errors.New("no create object function to decode")
	ErrNoRecordObjectToDecode = errors.New("no record object function to decode")
	ErrDrawVersion            = errors.New("draw version is invalid")
	ErrDayUnknown             = errors.New("day is unknown")
	ErrInvalidMoney           = errors.New("input money is invalid")
	ErrMoneyConverter         = errors.New("input money is not convertible")
	ErrInvalidReaderInput     = errors.New("could not use a nil reader")
	ErrInvalidFDJSource       = errors.New("fdj source doesn't match as expected the csv file")
)
