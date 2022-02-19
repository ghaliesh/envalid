package envalidator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetEnvFields(t *testing.T) {
	for _, tc := range getEnvFieldsTC {
		t.Run(tc.name, func(t *testing.T) {
			result := getEnvFields(tc.arg)

			require.Equal(t, len(tc.result), len(result))

			for i, v := range result {
				expected := tc.result[i]
				require.Equal(t, expected["type"], v["type"])
				require.Equal(t, expected["key"], v["key"])
			}
		})
	}
}
