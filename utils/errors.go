package utils

import "fmt"

func KeyDoesNotExistsError(key string) error {
	return fmt.Errorf("%s is missing from .env file", key)
}
