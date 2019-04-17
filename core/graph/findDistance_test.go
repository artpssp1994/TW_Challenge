package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindDistanceByRoute_SUCCESS(t *testing.T) {
	res1, err1 := FindDistanceByRoute(mockRailWays, "A-B-C")
	res2, err2 := FindDistanceByRoute(mockRailWays, "A-D")
	res3, err3 := FindDistanceByRoute(mockRailWays, "A-D-C")
	res4, err4 := FindDistanceByRoute(mockRailWays, "A-E-B-C-D")
	res5, err5 := FindDistanceByRoute(mockRailWays, "A-E-D")

	assert.Equal(t, "9", res1)
	assert.Nil(t, err1)
	assert.Equal(t, "5", res2)
	assert.Nil(t, err2)
	assert.Equal(t, "13", res3)
	assert.Nil(t, err3)
	assert.Equal(t, "22", res4)
	assert.Nil(t, err4)
	assert.Equal(t, "NO SUCH ROUTE", res5)
	assert.Nil(t, err5)
}

func Test_FindDistanceByRoute_SUCCESS_NotFoundCase(t *testing.T) {
	res, err := FindDistanceByRoute(mockRailWays, "A-B-X")

	assert.Equal(t, "NO SUCH ROUTE", res)
	assert.Nil(t, err)
}

func Test_FindDistanceByRoute_FAIL_EmptyCase(t *testing.T) {
	res, err := FindDistanceByRoute(mockRailWays, "")

	assert.Equal(t, "", res)
	assert.NotNil(t, err)
}
