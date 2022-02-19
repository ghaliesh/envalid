package utils

import "fmt"

func getRepeatedKeyError(keyname string) error {
	return fmt.Errorf("Duplicate key exist in .env file, duplicate key name: %s", keyname)
}

func getNotFoundKeyError(keyname string) error {
	return fmt.Errorf("Key doesn't exist, key name: %s", keyname)
}

func FindOne(m map[string]string, target string) (string, error) {
	var key string
	var count uint

	for k, _ := range m {
		if k == target {
			count += 1
			key = k
		}
	}

	if count > 1 {
		return "", getRepeatedKeyError(target)
	}

	if count > 0 {
		return key, nil
	}

	return "", getNotFoundKeyError(target)
}
