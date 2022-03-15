package envalid

import (
	"github.com/ghaliesh/envalid/envalidator"
)

type EnvalidInterface interface {
	ValidateEnv(rules interface{}, path string)
}

type OnErrorHandler = func(err error)

type Envalid struct {
	OnError OnErrorHandler
}

func (e *Envalid) ValidateEnv(validator interface{}, path string) {
	err := envalidator.Validate(validator, path)

	if err != nil {
		e.OnError(err)
	}
}
