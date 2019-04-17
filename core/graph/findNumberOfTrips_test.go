package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumberOfTrips_SUCCESS(t *testing.T) {
	res := FindNuberOfTripsWithStop(mockRailWays, "C", "C", 0, 3, false)
	res2 := FindNuberOfTripsWithStop(mockRailWays, "A", "C", 0, 4, true)
	assert.Equal(t, 2, res)
	assert.Equal(t, 3, res2)
}

func Test_NumberOfTrips_NotFoundCase(t *testing.T) {
	res := FindNuberOfTripsWithStop(mockRailWays, "C", "Z", 0, 10, false)
	res2 := FindNuberOfTripsWithStop(mockRailWays, "A", "Y", 0, 10, true)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, res2)
}

func Test_FindNuberOfTripsWithDistance_SUCCESS(t *testing.T) {
	res := FindNuberOfTripsWithDistance(mockRailWays, "C", "C", 0, 10)
	res2 := FindNuberOfTripsWithDistance(mockRailWays, "C", "C", 0, 30)
	assert.Equal(t, 1, res)
	assert.Equal(t, 7, res2)
}

func Test_FindNuberOfTripsWithDistance_NotFoudRoute(t *testing.T) {
	res := FindNuberOfTripsWithDistance(mockRailWays, "C", "G", 0, 10)
	res2 := FindNuberOfTripsWithDistance(mockRailWays, "C", "C", 0, 0)
	assert.Equal(t, 0, res)
	assert.Equal(t, 0, res2)
}
