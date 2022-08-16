package csvparser

import "github.com/pkg/errors"

// error for the csv parser
var (
	ErrOBJDecode              = errors.New("error to decode specific object")
	ErrNewDecoder             = errors.New("error to create a csvutil.NewDecoder")
	ErrBadConfiguration       = errors.New("configuration are not full setup to decode csv")
	ErrNoCreateObjectToDecode = errors.New("no create object function to decode")
	ErrRecorder               = errors.New("function record was return an error")
	ErrNilReader              = errors.New("reader is nil")
)
