package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/amodhakal/linear-reg-calculator/util"
)

func main() {
	filename := os.Args[1]

	orderedPairs := util.ReadFile(filename)
	xMean, yMean := util.CalculateSampleMeans(orderedPairs)
	intercept := util.CalculateIntecept(orderedPairs, xMean, yMean)
	slope := util.CalculateSlope(intercept, xMean, yMean)

	if len(os.Args) == 3 && os.Args[2] == "display" {
		util.PrintResults(len(orderedPairs), slope, intercept)
	} else if len(os.Args) == 5 && os.Args[2] == "calculate" && os.Args[3] == "y" {
		xVal, err := strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			log.Fatal("Number expected for the 4th argument")
		}

		fmt.Printf("Expected y-value: %.4f\n", (xVal*slope + intercept))

	} else if len(os.Args) == 5 && os.Args[2] == "calculate" && os.Args[3] == "x" {
		yVal, err := strconv.ParseFloat(os.Args[4], 64)
		if err != nil {
			log.Fatal("Number expected for the 4th argument")
		}

		fmt.Printf("Expected y-value: %.4f\n", ((yVal - intercept) / slope))

	} else {
		log.Fatal("Invalid arguments")
	}
}
