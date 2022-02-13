package envfilereader

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

var envReader = EnvFileReader{}

var _, b, _, _ = runtime.Caller(0)
var basepath = filepath.Dir(b)
var dotenvpath = fmt.Sprintf("%s/.env", basepath)

func createTmpFile(data []byte) error {
	err := os.WriteFile(dotenvpath, data, 0664)

	return err
}

func TestReadEnvFile(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			envFile := []byte(tc.envFileContent)
			err := createTmpFile(envFile)
			require.NoError(t, err)

			actualResult := envReader.ReadEnvFile(dotenvpath)
			expectedResult := tc.expectedResult
			require.Equal(t, expectedResult, actualResult)
		})
	}
}
