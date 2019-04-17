package graph

//FindNuberOfTripsWithStop - find possible route by stop limitation and can set specfic stop number by flag
func FindNuberOfTripsWithStop(railWays map[string]map[string]int, sStart string, sStop string, cStop int,
	maxStop int, isNeedExactlyStop bool) int {

	count := 0
	//stop recursive if it reach max stop
	if cStop > maxStop {
		return 0
	}

	//add possible ways to reach destination
	if sStart == sStop && cStop > 0 {
		if isNeedExactlyStop && cStop == maxStop {
			count += 1
		} else if !isNeedExactlyStop {
			count += 1
		}
	}

	cStop += +1
	for city, _ := range railWays[sStart] {
		n := FindNuberOfTripsWithStop(railWays, city, sStop, cStop, maxStop, isNeedExactlyStop)
		count += n
	}

	return count
}

//FindNuberOfTripsWithDistance - find possible route by distance limitation
func FindNuberOfTripsWithDistance(railWays map[string]map[string]int, sStart string, sStop string, cDistance int,
	maxDistance int) int {
	count := 0
	//stop recursive if it reach max distance
	if cDistance >= maxDistance {
		return 0
	}

	if sStart == sStop && cDistance > 0 {
		count += 1
	}

	for city, _ := range railWays[sStart] {
		d := cDistance + railWays[sStart][city]
		n := FindNuberOfTripsWithDistance(railWays, city, sStop, d, maxDistance)
		count += n
	}

	return count
}
