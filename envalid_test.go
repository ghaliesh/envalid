package envalid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var funcIsCalled = false

func onErrorHandler(err error) {
	funcIsCalled = true
}

func TestValidate(t *testing.T) {
	testSubject := Envalid{
		validator: struct{}{},
		OnError:   nil,
		path:      "/random/path",
	}
	wrapper := func() { testSubject.ValidateEnv() }
	require.Panics(t, wrapper)

	testSubject.path = ""
	testSubject.OnError = onErrorHandler
	testSubject.ValidateEnv()
	require.True(t, funcIsCalled)

}
