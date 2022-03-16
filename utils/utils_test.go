package utils

import (
	"os"
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
	resultsInt := []int64{}
	resultsUInt := []uint32{}
	resultsFloat := []float64{}

	result := RandomString(length)
	require.Len(t, result, length)

	results = append(results, result)

	for range timesToRun {
		res := RandomString(length)
		resInt := RandomInt()
		resFloat := RandomFloat()
		resUnit := RandomUnit()
		require.NotContains(t, results, res)
		require.NotContains(t, resultsInt, resInt)
		require.NotContains(t, resultsFloat, resFloat)
		require.NotContains(t, resultsUInt, resUnit)
		resultsUInt = append(resultsUInt, resUnit)
		resultsFloat = append(resultsFloat, resFloat)
		resultsInt = append(resultsInt, resInt)
		results = append(results, res)
	}
}

func TestExists(t *testing.T) {
	for _, tc := range existsTestcases {
		t.Run(tc.name, func(t *testing.T) {
			result := Exists(tc.dict, tc.target, tc.lookForKeys)
			require.Equal(t, tc.shouldExist, result)
		})
	}
}

func TestCreateFileTmp(t *testing.T) {
	data := RandomString(10)
	CreateTmpEnvFile([]byte(data))

	require.FileExists(t, "../tmp/.env")
	file, err := os.ReadFile("../tmp/.env")
	require.NoError(t, err)
	require.Equal(t, string(file), data)
}

func TestKeyDoesNotExistError(t *testing.T) {
	err := KeyDoesNotExistsError("key")
	require.Error(t, err)
	require.Equal(t, err.Error(), "envalid: key is missing from .env file")
}

func TestKeyKeyIsNotOfRightTypeError(t *testing.T) {
	err := KeyIsNotOfRightTypeError("key", "int", "value")
	require.Error(t, err)
}
