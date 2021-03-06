package reader

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghaliesh/envalid/utils"
)

const (
	linesep = "\n"
	pairsep = "="
)

type EnvKeyValuePairs = map[string]string

func getFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		panic("Can't read file in path")
	}

	return file
}

func stringify(file []byte) string {
	stringified := string(file)
	if len(stringified) == 0 {
		panic("File is either empty or invalid")
	}

	return stringified
}

func getKeypair(pair string) (key string, value string) {
	_keyval := strings.Split(pair, pairsep)
	keyval := utils.Filter(_keyval, func(s string) bool { return len(s) > 0 })
	len := len(keyval)

	// making sure the format is always [key, value]
	if len == 1 {
		keyval = []string{keyval[0], ""}
	} else if len > 2 {
		key := fmt.Sprintf("%s%s", keyval[0], pairsep)
		val := strings.Replace(pair, key, "", 1)
		keyval = []string{keyval[0], val}
	}

	key = keyval[0]
	value = keyval[1]

	return strings.TrimSpace(key), strings.TrimSpace(value)
}

func ReadEnvFile(path string) EnvKeyValuePairs {
	file := getFile(path)
	str := stringify(file)
	pairs := strings.Split(str, linesep)

	var result = make(EnvKeyValuePairs)

	for _, p := range pairs {
		if len(p) > 1 {
			key, value := getKeypair(strings.TrimSpace(p))
			if len(key) != 0 {
				result[key] = value
			}
		}
	}

	return result
}
