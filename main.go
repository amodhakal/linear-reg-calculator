package main

import "github.com/amodhakal/linear-reg-calculator/util"

func main() {
	orderedPairs := util.ReadFileFromArgs()

	// The below line is just there so I don't get an error for not using the variable
	println(len(orderedPairs))
}
