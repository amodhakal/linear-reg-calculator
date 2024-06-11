package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/amodhakal/linear-reg-calculator/util"
)

func main() {
	filename := os.Args[1]

	orderedPairs := util.ReadFile(filename)
	xMean, yMean := util.CalculateSampleMeans(orderedPairs)
	intercept := util.CalculateIntecept(orderedPairs, xMean, yMean)
	slope := util.CalculateSlope(intercept, xMean, yMean)
	reader := bufio.NewReader(os.Stdin)

	for {
		print("<cmd> ")
		line, _ := reader.ReadString('\n')

		if line == "display\n" {
			util.PrintResults(len(orderedPairs), slope, intercept)
			continue
		}

		if line == "quit\n" {
			break
		}

		cmd, direction, val := "", "", 0.0
		result, _ := fmt.Sscanf(line, "%s %s %f\n", &cmd, &direction, &val)

		if result != 3 || cmd != "calculate" {
			println("Invalid command")
		} else if direction == "y" {
			fmt.Printf("Expected y-value: %.4f\n", (val*slope + intercept))
		} else if direction == "x" {
			fmt.Printf("Expected x-value: %.4f\n", ((val - intercept) / slope))
		} else {
			println("Invalid command")
		}
	}
}
