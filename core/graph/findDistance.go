package graph

import (
	"errors"
	"strconv"
	"strings"
)

//FindDistanceByRoute - find distance by specific route
func FindDistanceByRoute(railWays map[string]map[string]int, path string) (string, error) {
	cityList := strings.Split(path, "-")
	result := ""

	if len(cityList) <= 1 {
		return "", errors.New("worng input format")
	}

	distance := 0
	startCity := cityList[0]

	for _, desCity := range cityList[1:] {
		if _, ok := railWays[startCity][desCity]; !ok {
			distance = -1
			break
		}

		distance = distance + railWays[startCity][desCity]
		startCity = desCity
	}

	if distance > 0 {
		result = strconv.Itoa(distance)
	} else {
		result = "NO SUCH ROUTE"
	}

	return result, nil
}
