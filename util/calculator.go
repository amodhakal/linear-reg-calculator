package util

func CalculateSampleMeans(orderedPairs [][2]float64) (float64, float64) {
	var xSum, ySum, count float64 = 0, 0, float64(len(orderedPairs))

	for _, orderedPair := range orderedPairs {
		xSum += orderedPair[0]
		ySum += orderedPair[1]
	}

	return xSum / count, ySum / count
}

func CalculateSlope(orderedPairs [][2]float64, xMean float64, yMean float64) float64 {
	var slope float64 = 0

	// TODO: Calculate slope

	return slope
}
