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

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Can't read the given file")
	}

	defer file.Close()
	data := make([][2]float64, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var xVal, yVal float64

		scanned, err := fmt.Sscanf(line, "%f, %f", &xVal, &yVal)
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
