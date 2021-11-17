package validation

import (
	"errors"
	"fmt"
)

var (
	ErrMissingFile         = errors.New("missing file")
	ErrNotAllowedFile      = errors.New("file not allowed")
	ErrInvalidJson         = errors.New("invalid json")
	ErrInvalidImgDimension = errors.New("invalid file dimension")

	ErrInvalidFileCase       = errors.New("invalid file name case")
	ErrInvalidFileExt        = errors.New("invalid file extension")
	ErrInvalidFileSize       = errors.New("invalid file size")
	ErrInvalidFileNameLength = errors.New("invalid file name length")
	ErrInvalidAddress        = errors.New("invalid address")
	//ErrInvalidFileNamePrefix  = errors.New("invalid file name prefix")
	//ErrInvalidFileNameFormat  = errors.New("invalid file name format")
	//ErrInvalidInfoAssetPaylod = errors.New("invalid info asset payload")
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
		msg = fmt.Sprintf("%s - validation error: %s \n", msg, err.Error())
	}

	return msg
}

func (e *ErrComposite) Append(err error) {
	e.errors = append(e.errors, err)
}

func (e *ErrComposite) GetErrors() []error {
	return e.errors
}

type Warning struct {
	err error
}

func NewWarning(err error) *Warning {
	return &Warning{err: err}
}

func (e *Warning) Error() string {
	return fmt.Sprintf("warrning: %s", e.err.Error())
}
