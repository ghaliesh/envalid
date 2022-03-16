package envalid

import (
	"log"

	"github.com/ghaliesh/envalid/envalidator"
	"github.com/ghaliesh/envalid/utils"
)

type EnvalidInterface interface {
	ValidateEnv(rules interface{}, path string)
}

type OnErrorHandler = func(err error)

type Envalid struct {
	OnError   OnErrorHandler
	validator interface{}
	path      string
}

func (e *Envalid) ValidateEnv() {
	if e.OnError == nil {
		log.Panic(utils.ErrOnErrorFunctionIsMissing)
	}

	if len(e.path) == 0 {
		e.OnError(utils.ErrPathIsNotProvided)
	}

	if e.validator == nil {
		e.OnError(utils.ErrValidatorIsNotProvided)
	}

	err := envalidator.Validate(e.validator, e.path)

	if err != nil {
		e.OnError(err)
	}
}
