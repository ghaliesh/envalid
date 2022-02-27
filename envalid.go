package envalid

import (
	"github.com/ghaliesh/envalid/envalidator"
)

type EnvalidInterface interface {
	ValidateEnv(rules interface{}, path string)
}

type Envalid struct{}

func (e *Envalid) ValidateEnv(validator interface{}, path string) {
	envalidator.Validate(validator, path)
}
