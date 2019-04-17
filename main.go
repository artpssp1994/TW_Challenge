package main

import (
	"ThoughWorkAssignMent/core/format"
	"ThoughWorkAssignMent/core/graph"
	"fmt"
	"os"
)

//main - Receive input call logic function and print result to console
func main() {

	if len(os.Args) < 2 {
		fmt.Println("Parameter is require!!!")
		return
	}

	inputGraph := os.Args[1]

	//Parse graph string from input to map
	mRailWays, err := format.ParseInputToMap(inputGraph)
	if err != nil {
		fmt.Printf("Error parseInputToMap: %+v\n", err)
		return
	}

	res1, err := graph.FindDistanceByRoute(mRailWays, "A-B-C")
	res2, err := graph.FindDistanceByRoute(mRailWays, "A-D")
	res3, err := graph.FindDistanceByRoute(mRailWays, "A-D-C")
	res4, err := graph.FindDistanceByRoute(mRailWays, "A-E-B-C-D")
	res5, err := graph.FindDistanceByRoute(mRailWays, "A-E-D")
	res6 := graph.FindNuberOfTripsWithStop(mRailWays, "C", "C", 0, 3, false)
	res7 := graph.FindNuberOfTripsWithStop(mRailWays, "A", "C", 0, 4, true)
	res8, _ := graph.FindShortedPath(mRailWays, "A", "C", 0, []string{})
	res9, _ := graph.FindShortedPath(mRailWays, "B", "B", 0, []string{})
	res10 := graph.FindNuberOfTripsWithDistance(mRailWays, "C", "C", 0, 30)

	//Handle error case input wrong format
	if err != nil {
		fmt.Printf("Error : %+v", err)
		return
	}

	//Print result to console
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
