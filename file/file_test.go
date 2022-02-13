package envfilereader

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var envReader = EnvFileReader{}
var dirname, _ = os.Getwd()
var path = filepath.Join(dirname, "../tmp/.env")
var invalidpath = "in/valid/pa/th"

func createTmpFile(data []byte) error {

	err := os.WriteFile(path, data, 0777)

	return err
}

func funcToPanicDuetoEmptyFile() {
	envReader.ReadEnvFile(path)
}

func funcToPanicDuetoNonExistentPath() {
	envReader.ReadEnvFile(invalidpath)
}

func TestReadEnvFile(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Create .env file
			envFile := []byte(tc.envFileContent)
			err := createTmpFile(envFile)
			require.NoError(t, err)

			// Edge case: Fucntion Panic
			if len(tc.envFileContent) == 0 {
				require.Panics(t, funcToPanicDuetoEmptyFile)
				return
			}

			// Test result
			actualResult := envReader.ReadEnvFile(path)
			expectedResult := tc.expectedResult
			require.Equal(t, expectedResult, actualResult)

		})
	}

	require.Panics(t, funcToPanicDuetoNonExistentPath)
}
