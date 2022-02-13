package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var sample string = RandomString(6)

var strs = []string{
	RandomString(3),
	RandomString(4),
	RandomString(5),
}

func removeSample(s string) bool { return s != sample }

func removeNothing(s string) bool { return true }

func removeEverything(s string) bool { return false }

func TestFilter(t *testing.T) {
	// expectedInex = len(strs)
	strsWithSample := append(strs, sample)

	strsWithoutSample := Filter(strsWithSample, removeSample)
	require.Equal(t, strs, strsWithoutSample)

	strsAfterRemovingNothing := Filter(strs, removeNothing)
	require.Equal(t, strs, strsAfterRemovingNothing)

	strsAfterRemovingEverything := Filter(strs, removeEverything)
	require.Zero(t, len(strsAfterRemovingEverything))

}
