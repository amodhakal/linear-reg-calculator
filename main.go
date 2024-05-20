package main

import "github.com/amodhakal/linear-reg-calculator/util"

func main() {
	orderedPairs := util.ReadFileFromArgs()
	xMean, yMean := util.CalculateSampleMeans(orderedPairs)
	slope := util.CalculateSlope(orderedPairs, xMean, yMean)

	var intecept, rSquared float64 = 0, 0
	util.PrintResults(len(orderedPairs), slope, float64(intecept), float64(rSquared))
}
