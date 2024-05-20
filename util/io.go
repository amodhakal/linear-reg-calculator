package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFileFromArgs() [][2]float64 {
	if len(os.Args) != 2 {
		log.Fatal("Invalid arguments")
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Can't read the given file")
	}

	defer file.Close()
	data := make([][2]float64, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var xVal, yVal float64

		scanned, err := fmt.Sscanf(line, "%f,%f", &xVal, &yVal)
		if err != nil {
			println(err.Error())
		}

		if scanned != 2 {
			log.Fatalf("Error around line: %d!", len(data))
		}

		pair := [2]float64{xVal, yVal}
		data = append(data, pair)
	}

	if scanner.Err() != nil {
		panic("Can't read the given file")
	}

	return data
}

func PrintResults(count int, slope float64, intercept float64, rSquared float64) {
	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Slope: %.4f\n", slope)
	fmt.Printf("Intercept: %.4f\n", intercept)
	fmt.Printf("R2: %.4f\n", rSquared)
}
