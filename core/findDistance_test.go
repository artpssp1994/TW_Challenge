package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockRailWays = map[string]map[string]int{
	"A": map[string]int{
		"B": 5,
		"D": 5,
		"E": 7,
	},
	"B": map[string]int{
		"C": 4,
	},
	"C": map[string]int{
		"D": 8,
		"E": 2,
	},
	"D": map[string]int{
		"C": 8,
		"E": 6,
	},
	"E": map[string]int{
		"B": 3,
	},
}

func Test_findDistanceByRoute_SUCCESS(t *testing.T) {
	res1, err1 := findDistanceByRoute(mockRailWays, "A-B-C")
	res2, err2 := findDistanceByRoute(mockRailWays, "A-D")
	res3, err3 := findDistanceByRoute(mockRailWays, "A-D-C")
	res4, err4 := findDistanceByRoute(mockRailWays, "A-E-B-C-D")
	res5, err5 := findDistanceByRoute(mockRailWays, "A-E-D")

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

func Test_findDistanceByRoute_SUCCESS_NotFoundCase(t *testing.T) {
	res, err := findDistanceByRoute(mockRailWays, "A-B-X")

	assert.Equal(t, "NO SUCH ROUTE", res)
	assert.Nil(t, err)
}

func Test_findDistanceByRoute_FAIL_EmptyCase(t *testing.T) {
	res, err := findDistanceByRoute(mockRailWays, "")

	assert.Equal(t, "", res)
	assert.NotNil(t, err)
}
