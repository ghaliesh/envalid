package envalidator

import (
	"testing"

	"github.com/ghaliesh/envalid/utils"
	"github.com/stretchr/testify/require"
)

func TestCheckExists(t *testing.T) {
	for _, tc := range CheckExistTestCases {
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
