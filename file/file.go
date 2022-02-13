package envfilereader

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

type Reader interface {
	ReadEnvFile(path string) string
}

type EnvFileReader struct{}

type EnvKeyValuePairs = map[string]string

func (ef *EnvFileReader) getFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		panic("Can't read file in path")
	}

	return file
}

func (ef *EnvFileReader) stringify(file []byte) string {
	stringified := string(file)
	if len(stringified) == 0 {
		panic("File is either empty or invalid")
	}

	return stringified
}

func (ef *EnvFileReader) getKeypair(pair string) (key string, value string) {
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

func (ef *EnvFileReader) ReadEnvFile(path string) EnvKeyValuePairs {
	file := ef.getFile(path)
	str := ef.stringify(file)
	pairs := strings.Split(str, linesep)

	var result = make(EnvKeyValuePairs)

	for _, p := range pairs {
		if len(p) > 1 {
			key, value := ef.getKeypair(strings.TrimSpace(p))
			if len(key) != 0 {
				result[key] = value
			}
		}
	}

	return result
}
