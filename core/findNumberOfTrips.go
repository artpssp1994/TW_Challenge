package core

func findNuberOfTrips(railWays map[string]map[string]int, sStart string, sStop string, cStop int,
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
	for k, _ := range railWays[sStart] {
		c := findNuberOfTrips(railWays, k, sStop, cStop, maxStop, isNeedExactlyStop)
		count += c
	}

	return count
}
