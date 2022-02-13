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

var (
	length     = 5
	timesToRun = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func TestRandom(t *testing.T) {
	results := []string{}

	result := RandomString(length)
	require.Len(t, result, length)

	results = append(results, result)

	for range timesToRun {
		res := RandomString(length)
		require.NotContains(t, results, res)

		results = append(results, res)
	}
}
