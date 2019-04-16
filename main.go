package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	_, err := parseInputToMap("")
	if err != nil {
		fmt.Printf("Error parseInputToMap: %+v\n", err)
		return
	}

}

func parseInputToMap(input string) (map[string]map[string]int, error) {
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

		if _, ok := m[startCity]; !ok {
			m[startCity] = make(map[string]int)
		}

		m[startCity][destinationCity] = length
	}

	return m, nil
}
