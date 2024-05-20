package main

import "github.com/amodhakal/linear-reg-calculator/util"

func main() {
	orderedPairs := util.ReadFileFromArgs()
	xMean, yMean := util.CalculateSampleMeans(orderedPairs)
	intercept := util.CalculateIntecept(orderedPairs, xMean, yMean)
	slope := util.CalculateSlop(intercept, xMean, yMean)

	var rSquared float64 = 0
	util.PrintResults(len(orderedPairs), slope, intercept, rSquared)
}
