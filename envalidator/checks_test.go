package envalidator

import (
	"testing"

	"github.com/ghaliesh/envalid/utils"
	"github.com/stretchr/testify/require"
)

func TestCheckExists(t *testing.T) {
	for _, tc := range checkExistTestCases {
		t.Run(tc.name, func(t *testing.T) {
			err := checkKeyExist(tc.keyValPairs, tc.key)

			if tc.shouldPanic {
				expectedError := utils.KeyDoesNotExistsError(tc.key)
				require.EqualError(t, err, expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestChecks(t *testing.T) {
	for _, main := range checksTc {
		for _, tc := range main.checkFuncsData {

			t.Run(tc.name, func(t *testing.T) {

				err := tc.funcInTest(main.arg, "key")
				if tc.shouldPanic {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
				}
			})
		}
	}
}

func TestChecksType(t *testing.T) {
	for _, main := range checkTypeTcs {
		for _, typeof := range main.matchingTypes {

			t.Run(main.name, func(t *testing.T) {

				err := checkType(typeof, main.key, main.value)
				require.NoError(t, err)
			})
		}

		for _, types := range main.nonMatchingTyps {
			for _, typeof := range types {

				t.Run(main.name, func(t *testing.T) {

					err := checkType(typeof, main.key, main.value)
					require.Error(t, err)
				})
			}
		}

	}
}
