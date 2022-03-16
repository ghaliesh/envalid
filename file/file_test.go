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

			// Test result
			actualResult, err := ReadEnvFile(path)
			expectedResult := tc.expectedResult
			if len(tc.envFileContent) == 0 {
				require.EqualError(t, err, utils.ErrInvalidEnvFile.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, expectedResult, actualResult)
			}
		})
	}

	var invalidpath = "in/valid/pa/th"
	_, err := ReadEnvFile(invalidpath)
	require.Error(t, err)

}
