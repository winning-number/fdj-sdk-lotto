package reader

import "github.com/pkg/errors"

// error list
var (
	ErrInvalidReaderInput = errors.New("could not use a nil reader")
	ErrInvalidArchive     = errors.New("invalid number files inside the archive")
)
