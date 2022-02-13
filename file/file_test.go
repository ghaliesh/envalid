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

func createTmpFile(data []byte) error {

	err := os.WriteFile(path, data, 0777)

	return err
}

func TestReadEnvFile(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Create .env file
			envFile := []byte(tc.envFileContent)
			err := createTmpFile(envFile)
			require.NoError(t, err)

			// Test result
			actualResult := envReader.ReadEnvFile(path)
			expectedResult := tc.expectedResult
			require.Equal(t, expectedResult, actualResult)

			// Clean up
			err = os.Remove(path)
			require.NoError(t, err)
		})
	}
}
