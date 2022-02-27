package envalidator

import (
	"testing"

	"github.com/ghaliesh/envalid/utils"
	"github.com/stretchr/testify/require"
)

func TestCheckExists(t *testing.T) {
	for _, tc := range checkExistTestCases {
		t.Run(tc.name, func(t *testing.T) {
			wrapper := func() { checkKeyExist(tc.keyValPairs, tc.key) }

			if tc.shouldPanic {
				err := utils.KeyDoesNotExistsError(tc.key)
				require.PanicsWithError(t, err.Error(), wrapper)
			} else {
				require.NotPanics(t, wrapper)
			}
		})
	}
}

func TestPanicIfError(t *testing.T) {
	for _, tc := range panicIfErrorTCs {
		t.Run(tc.name, func(t *testing.T) {
			wrapper := func() { panicIfError(tc.err, tc.key, tc.val, tc.val) }

			if tc.shouldPanic {
				require.Panics(t, wrapper)
			} else {
				require.NotPanics(t, wrapper)
			}
		})
	}
}

func TestChecks(t *testing.T) {
	for _, main := range checksTc {
		for _, tc := range main.checkFuncsData {

			t.Run(tc.name, func(t *testing.T) {

				wrapper := func() { tc.funcInTest(main.arg, "key") }
				if tc.shouldPanic {
					require.Panics(t, wrapper)
				} else {
					require.NotPanics(t, wrapper)
				}
			})
		}
	}
}

func TestChecksType(t *testing.T) {
	for _, main := range checkTypeTcs {
		for _, typeof := range main.matchingTypes {

			t.Run(main.name, func(t *testing.T) {

				wrapper := func() { checkType(typeof, main.key, main.value) }
				require.NotPanics(t, wrapper)
			})
		}

		for _, types := range main.nonMatchingTyps {
			for _, typeof := range types {

				t.Run(main.name, func(t *testing.T) {

					wrapper := func() { checkType(typeof, main.key, main.value) }
					require.Panics(t, wrapper)
				})
			}
		}

	}
}
