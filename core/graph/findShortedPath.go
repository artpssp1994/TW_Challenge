package graph

//FindShortedPath - find shorted path between two city
func FindShortedPath(railWays map[string]map[string]int, sStart string, sStop string,
	distance int, passedCity []string) (d int, isReached bool) {

	isReached = false
	//stop and return distance when reach destination
	if sStart == sStop && distance > 0 {
		return distance, true
	}

	minDistance := -1
	for city, _ := range railWays[sStart] {

		//check this city already pass or not
		if contains(passedCity, city) {
			continue
		}

		passedCity = append(passedCity, city)

		//copy slice for send to recursive to avoid slice point same address problem
		tmpPassedCity := make([]string, len(passedCity))
		copy(tmpPassedCity, passedCity)

		d := distance + railWays[sStart][city]
		sumDistance, iR := FindShortedPath(railWays, city, sStop, d, tmpPassedCity)
		if (minDistance == -1 || sumDistance < minDistance) && iR {
			minDistance = sumDistance
			isReached = iR
		}
	}

	if minDistance == -1 {
		minDistance = 0
	}

	return minDistance, isReached
}

//contains - check slice contain value or not
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
