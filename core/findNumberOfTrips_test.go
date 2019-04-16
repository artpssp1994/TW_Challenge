package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberOfTrips_SUCCESS(t *testing.T) {
	res := findNuberOfTrips(mockRailWays, "C", "C", 0, 3, false)
	res2 := findNuberOfTrips(mockRailWays, "A", "C", 0, 4, true)
	assert.Equal(t, 2, res)
	assert.Equal(t, 3, res2)
}

func Test_NumberOfTrips_NotFoundCase(t *testing.T) {
	res := findNuberOfTrips(mockRailWays, "C", "Z", 0, 10, false)
	res2 := findNuberOfTrips(mockRailWays, "A", "Y", 0, 10, true)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, res2)
}
