package envalidator

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ghaliesh/envalid/utils"
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

var dirname, _ = os.Getwd()
var path = filepath.Join(dirname, "../tmp/.env")

func TestValidate(t *testing.T) {
	for _, main := range validateTCs {
		for i, v := range main.validators {
			t.Run(main.name+"#"+fmt.Sprint(i), func(t *testing.T) {
				utils.CreateTmpEnvFile([]byte(main.envFile))
				err := Validate(v, path)

				if main.shouldPanic {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
				}
			})
		}
	}
}
