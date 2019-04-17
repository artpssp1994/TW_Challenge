package main

import (
	"ThoughWorkAssignMent/core"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Parameter is require!!!")
		return
	}

	inputGraph := os.Args[1]

	// "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
	mRailWays, err := parseInputToMap(inputGraph)
	if err != nil {
		fmt.Printf("Error parseInputToMap: %+v\n", err)
		return
	}

	res1, err := core.FindDistanceByRoute(mRailWays, "A-B-C")
	res2, err := core.FindDistanceByRoute(mRailWays, "A-D")
	res3, err := core.FindDistanceByRoute(mRailWays, "A-D-C")
	res4, err := core.FindDistanceByRoute(mRailWays, "A-E-B-C-D")
	res5, err := core.FindDistanceByRoute(mRailWays, "A-E-D")
	res6 := core.FindNuberOfTripsWithStop(mRailWays, "C", "C", 0, 3, false)
	res7 := core.FindNuberOfTripsWithStop(mRailWays, "A", "C", 0, 4, true)
	res8, _ := core.FindShortedPath(mRailWays, "A", "C", 0, []string{})
	res9, _ := core.FindShortedPath(mRailWays, "B", "B", 0, []string{})
	res10 := core.FindNuberOfTripsWithDistance(mRailWays, "C", "C", 0, 30)

	if err != nil {
		fmt.Printf("Error : %+v", err)
		return
	}

	fmt.Printf(
		"Output #1:  %+v\n"+
			"Output #2:  %+v\n"+
			"Output #3:  %+v\n"+
			"Output #4:  %+v\n"+
			"Output #5:  %+v\n"+
			"Output #6:  %+v\n"+
			"Output #7:  %+v\n"+
			"Output #8:  %+v\n"+
			"Output #9:  %+v\n"+
			"Output #10: %+v\n",
		res1, res2, res3, res4, res5, res6, res7, res8, res9, res10)

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
