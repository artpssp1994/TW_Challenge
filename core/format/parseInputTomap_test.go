package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInputToMap_SUCCESS(t *testing.T) {
	eValue := map[string]map[string]int{
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

	aValue, err := ParseInputToMap("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7")

	assert.Equal(t, eValue, aValue)
	assert.Nil(t, err)
}

func Test_parseInputToMap_FAIL(t *testing.T) {
	aValue, err := ParseInputToMap("AB5, BCW")

	assert.NotNil(t, err)
	assert.Nil(t, aValue)
}
