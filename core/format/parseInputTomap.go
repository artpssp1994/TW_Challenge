package format

import (
	"errors"
	"strconv"
	"strings"
)

//ParseInputToMap - Receive string graph and parse to map
//Example input : "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
func ParseInputToMap(input string) (map[string]map[string]int, error) {

	//Create two dimension map
	m := make(map[string]map[string]int)

	stringSlice := strings.Split(input, ", ")

	for _, v := range stringSlice {

		rRailway := []rune(v)
		if len(rRailway) != 3 {
			return nil, errors.New("worng input format")
		}

		startCity := string(rRailway[0])
		destinationCity := string(rRailway[1])
		length, errConver := strconv.Atoi(string(rRailway[2]))
		if errConver != nil {
			return nil, errConver
		}

		//Create new map key if not exist before
		if _, ok := m[startCity]; !ok {
			m[startCity] = make(map[string]int)
		}

		//Set distance to map by origination and destination
		m[startCity][destinationCity] = length
	}

	return m, nil
}
