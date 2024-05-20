package util

import "math"

func CalculateSampleMeans(orderedPairs [][2]float64) (float64, float64) {
	var xSum, ySum, count float64 = 0, 0, float64(len(orderedPairs))

	for _, orderedPair := range orderedPairs {
		xSum += orderedPair[0]
		ySum += orderedPair[1]
	}

	return xSum / count, ySum / count
}

func CalculateIntecept(orderedPairs [][2]float64, xMean float64, yMean float64) float64 {
	var numeratorSum, denominatorSum float64 = 0, 0

	for _, orderedPair := range orderedPairs {
		xMeanDiff := orderedPair[0] - xMean
		yMeanDiff := orderedPair[1] - yMean

		numeratorSum += xMeanDiff * yMeanDiff
		denominatorSum += math.Pow(xMeanDiff, 2)
	}

	return numeratorSum / denominatorSum
}

func CalculateSlop(intercept float64, xMean float64, yMean float64) float64 {
	return yMean - intercept*xMean
}