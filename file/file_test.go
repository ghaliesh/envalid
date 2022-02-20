package reader

import (
	"testing"

	"github.com/ghaliesh/envalid/utils"
	"github.com/stretchr/testify/require"
)

func TestReadEnvFile(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Create .env file
			envFile := []byte(tc.envFileContent)
			err := utils.CreateTmpEnvFile(envFile)
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
