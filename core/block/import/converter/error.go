package converter

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
)

var ErrCancel = fmt.Errorf("import is canceled")
var ErrNoObjectsToImport = fmt.Errorf("source path doesn't contain objects to import")

type ConvertError map[string]error

func NewError() ConvertError {
	return ConvertError{}
}

func NewFromError(name string, initialError error) ConvertError {
	ce := ConvertError{}

	ce.Add(name, initialError)

	return ce
}

func NewCancelError(path string, err error) ConvertError {
	wrappedError := errors.Wrap(ErrCancel, err.Error())
	cancelError := NewFromError(path, wrappedError)
	return cancelError
}

func (ce ConvertError) Add(objectName string, err error) {
	ce[objectName] = err
}

func (ce ConvertError) Merge(err ConvertError) {
	for fileName, errPb := range err {
		ce[fileName] = errPb
	}
}

func (ce ConvertError) IsEmpty() bool {
	return len(ce) == 0
}

func (ce ConvertError) Error() error {
	var pattern = "source: %s, error: %s" + "\n"
	var errorString bytes.Buffer
	if ce.IsEmpty() {
		return nil
	}
	for name, err := range ce {
		errorString.WriteString(fmt.Sprintf(pattern, name, err.Error()))
	}
	return fmt.Errorf(errorString.String())
}

func (ce ConvertError) Get(objectName string) error {
	return ce[objectName]
}
