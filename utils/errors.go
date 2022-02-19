package utils

import "fmt"

func KeyDoesNotExistsError(key string) error {
	return fmt.Errorf("%s is missing from .env file", key)
}

func KeyIsNotOfRightTypeError(key string, val string, typeof string) error {
	return fmt.Errorf("\n expected value of '%s' to be type of '%s' \n got value '%s'", key, typeof, val)
}
