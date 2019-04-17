package core

func findNuberOfTripsWithStop(railWays map[string]map[string]int, sStart string, sStop string, cStop int,
	maxStop int, isNeedExactlyStop bool) int {

	count := 0
	if cStop > maxStop {
		return 0
	}

	if sStart == sStop && cStop > 0 {
		if isNeedExactlyStop && cStop == maxStop {
			count += 1
		} else if !isNeedExactlyStop {
			count += 1
		}
	}

	cStop += +1
	for city, _ := range railWays[sStart] {
		n := findNuberOfTripsWithStop(railWays, city, sStop, cStop, maxStop, isNeedExactlyStop)
		count += n
	}

	return count
}

func findNuberOfTripsWithDistance(railWays map[string]map[string]int, sStart string, sStop string, cDistance int,
	maxDistance int) int {
	count := 0
	if cDistance >= maxDistance {
		return 0
	}

	if sStart == sStop && cDistance > 0 {
		count += 1
	}

	for city, _ := range railWays[sStart] {
		d := cDistance + railWays[sStart][city]
		n := findNuberOfTripsWithDistance(railWays, city, sStop, d, maxDistance)
		count += n
	}

	return count
}
