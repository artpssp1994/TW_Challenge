package core

func FindShortedPath(railWays map[string]map[string]int, sStart string, sStop string,
	distance int, passedCity []string) (d int, isReached bool) {

	isReached = false
	if sStart == sStop && distance > 0 {
		return distance, true
	}

	minDistance := -1
	for city, _ := range railWays[sStart] {

		if contains(passedCity, city) {
			continue
		}

		passedCity = append(passedCity, city)
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
