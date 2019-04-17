package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindShortedPath_SUCCESS(t *testing.T) {
	res, isFound := FindShortedPath(mockRailWays, "A", "C", 0, []string{})
	res2, isFound2 := FindShortedPath(mockRailWays, "B", "B", 0, []string{})

	assert.Equal(t, 9, res)
	assert.Equal(t, true, isFound)
	assert.Equal(t, 9, res2)
	assert.Equal(t, true, isFound2)
}

func Test_FindShortedPath_NotFound(t *testing.T) {
	res, isFound := FindShortedPath(mockRailWays, "A", "Z", 0, []string{})
	res2, isFound2 := FindShortedPath(mockRailWays, "B", "B", 0, []string{"A", "B", "C", "D"})

	assert.Equal(t, 0, res)
	assert.Equal(t, false, isFound)
	assert.Equal(t, 0, res2)
	assert.Equal(t, false, isFound2)
}

func Test_contain(t *testing.T) {
	passedCity := []string{"A", "B", "C", "D"}
	res := contains(passedCity, "A")
	res2 := contains(passedCity, "E")

	assert.Equal(t, true, res)
	assert.Equal(t, false, res2)
}
