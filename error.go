package lotto

import "github.com/pkg/errors"

// error for the csv parser
var (
	ErrInvalidArchive         = errors.New("invalid number files inside the archive")
	ErrDrawTypeKeyNotFound    = errors.New("draw type key not found in the context recorder")
	ErrDrawTypeDecode         = errors.New("wrong draw type to decode")
	ErrNoCreateObjectToDecode = errors.New("no create object function to decode")
	ErrNoRecordObjectToDecode = errors.New("no record object function to decode")
	ErrDrawVersion            = errors.New("draw version is invalid")
)
