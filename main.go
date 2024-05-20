package main

import "github.com/amodhakal/linear-reg-calculator/util"

func main() {
	orderedPairs := util.ReadFileFromArgs()
	var slope, intecept, rSquared float64 = 0, 0, 0

	util.PrintResults(len(orderedPairs), slope, float64(intecept), float64(rSquared))
}
