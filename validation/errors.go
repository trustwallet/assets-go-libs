package validation

import (
	"errors"
	"fmt"
)

var (
	ErrMissingFile           = errors.New("missing file")
	ErrMissingField          = errors.New("missing required fields")
	ErrInvalidField          = errors.New("invalid field")
	ErrNotAllowedFile        = errors.New("file is not allowed")
	ErrInvalidAddress        = errors.New("invalid address")
	ErrInvalidJson           = errors.New("invalid json")
	ErrInvalidImgDimension   = errors.New("invalid file dimension")
	ErrInvalidFileNameCase   = errors.New("invalid file name case")
	ErrInvalidFileExt        = errors.New("invalid file extension")
	ErrInvalidFileSize       = errors.New("invalid file size")
	ErrInvalidFileNameLength = errors.New("invalid file name length")
)

func NewErrComposite() *ErrComposite {
	return &ErrComposite{}
}

type ErrComposite struct {
	errors []error
}

func (e *ErrComposite) Len() int {
	return len(e.errors)
}

func (e *ErrComposite) Error() string {
	var msg string
	for _, err := range e.errors {
		msg += fmt.Sprintf("- %s\n", err.Error())
	}

	return msg
}

func (e *ErrComposite) Append(err error) {
	e.errors = append(e.errors, err)
}

func (e *ErrComposite) GetErrors() []error {
	return e.errors
}
