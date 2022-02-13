package envfilereader

import (
	"fmt"
	"os"
	"strings"
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
	keyval := strings.Split(pair, pairsep)
	len := len(keyval)

	// making sure the format is always [key, value]
	if len == 1 {
		keyval = []string{keyval[0], ""}
	} else if len > 2 {
		key := fmt.Sprintf("%s%s", keyval[0], pairsep)
		val := strings.Replace(pair, key, "", 1)
		keyval = []string{keyval[0], val}
	} else if len == 0 {
		keyval = []string{"", ""}
	}

	key = keyval[0]
	value = keyval[1]

	return strings.TrimSpace(key), strings.TrimSpace(value)
}

func (ef *EnvFileReader) ReadEnvFile(path string) EnvKeyValuePairs {
	file := ef.getFile(path)
	str := ef.stringify(file)

	pairs := strings.Split(str, linesep)
	if len(pairs) == 0 {
		panic(".env file is invalid")
	}

	var result = make(EnvKeyValuePairs)

	for _, p := range pairs {
		if len(p) > 1 {
			key, value := ef.getKeypair(p)
			if len(key) != 0 {
				result[key] = value
			}
		}
	}

	return result

}
