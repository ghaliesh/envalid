package utils

import (
	"errors"
	"fmt"
)

// user error
var ErrOnErrorFunctionIsMissing = errors.New("envalid: onError is required")
var ErrValidatorIsNotProvided = errors.New("envalid: validator is required")
var ErrPathIsNotProvided = errors.New("envalid: path is required")
var ErrInvalidEnvFile = errors.New("envalid: env file is either empty or invalid")

// internal error
var ErrUnhandledTypes = errors.New("envalid: unknown types, only primitive types are allowed")

func KeyDoesNotExistsError(key string) error {
	return fmt.Errorf("envalid: %s is missing from .env file", key)
}

func KeyIsNotOfRightTypeError(key string, val string, typeof string) error {
	return fmt.Errorf("envalid: \n expected value of '%s' to be type of '%s' \n got value '%s'", key, typeof, val)
}
