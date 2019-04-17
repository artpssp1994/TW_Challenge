package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberOfTrips_SUCCESS(t *testing.T) {
	res := findNuberOfTripsWithStop(mockRailWays, "C", "C", 0, 3, false)
	res2 := findNuberOfTripsWithStop(mockRailWays, "A", "C", 0, 4, true)
	assert.Equal(t, 2, res)
	assert.Equal(t, 3, res2)
}

func Test_NumberOfTrips_NotFoundCase(t *testing.T) {
	res := findNuberOfTripsWithStop(mockRailWays, "C", "Z", 0, 10, false)
	res2 := findNuberOfTripsWithStop(mockRailWays, "A", "Y", 0, 10, true)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, res2)
}

func Test_findNuberOfTripsWithDistance_SUCCESS(t *testing.T) {
	res := findNuberOfTripsWithDistance(mockRailWays, "C", "C", 0, 10)
	res2 := findNuberOfTripsWithDistance(mockRailWays, "C", "C", 0, 30)
	assert.Equal(t, 1, res)
	assert.Equal(t, 7, res2)
}

func Test_findNuberOfTripsWithDistance_NotFoudRoute(t *testing.T) {
	res := findNuberOfTripsWithDistance(mockRailWays, "C", "G", 0, 10)
	res2 := findNuberOfTripsWithDistance(mockRailWays, "C", "C", 0, 0)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, res2)
}
