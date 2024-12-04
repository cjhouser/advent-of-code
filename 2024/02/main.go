package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Find the distance between two numbers
func distance(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	// Read the data
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	safe_count := 0
	for {
		report, err := reader.Read()
		if err == io.EOF {
			break
		}

		allInRange := true
		sum := 0

		for level := 1; level < len(report); level++ {
			prev, _ := strconv.Atoi(report[level-1])
			curr, _ := strconv.Atoi(report[level])
			if prev < curr {
				sum--
			} else if prev > curr {
				sum++
			}

			dist := distance(prev, curr)
			allInRange = allInRange && (dist >= 1 && dist <= 3)
		}

		if allInRange && distance(0, sum) == len(report)-1 {
			safe_count++
		}
	}

	fmt.Printf("safe reports: %d\n",
		safe_count,
	)
}
