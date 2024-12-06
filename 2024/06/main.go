package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Position struct {
	Blocked bool
	Left    bool
	Right   bool
	Up      bool
	Down    bool
}

func main() {
	// Read the data
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [][]Position

	var guardRow int
	var guardCol int
	var guardFacing int

	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []Position
		for col, char := range line {
			if char == '^' {
				guardRow = rowCount
				guardCol = col
				guardFacing = 0
				row = append(row, Position{Blocked: false, Left: false, Right: false, Up: false, Down: false})
			} else if char == '>' {
				guardRow = rowCount
				guardCol = col
				guardFacing = 1
				row = append(row, Position{Blocked: false, Left: false, Right: false, Up: false, Down: false})
			} else if char == 'v' {
				guardRow = rowCount
				guardCol = col
				guardFacing = 2
				row = append(row, Position{Blocked: false, Left: false, Right: false, Up: false, Down: false})
			} else if char == '<' {
				guardRow = rowCount
				guardCol = col
				guardFacing = 3
				row = append(row, Position{Blocked: false, Left: false, Right: false, Up: false, Down: false})
			} else if char == '#' {
				row = append(row, Position{Blocked: true, Left: false, Right: false, Up: false, Down: false})
			} else {
				row = append(row, Position{Blocked: false, Left: false, Right: false, Up: false, Down: false})
			}
		}
		matrix = append(matrix, row)
		rowCount++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	visitedCount := 0
	for {
		// Guard has left the map
		if guardRow < 0 || guardRow >= len(matrix) || guardCol < 0 || guardCol >= len(matrix[0]) {
			break
		}

		// Distinct locations count toward result
		if !matrix[guardRow][guardCol].Up && !matrix[guardRow][guardCol].Right && !matrix[guardRow][guardCol].Down && !matrix[guardRow][guardCol].Left {
			visitedCount++
		}

		// Track current position and direction, then move the guard
		if guardFacing == 0 {
			matrix[guardRow][guardCol].Up = true
			guardRow--
		} else if guardFacing == 1 {
			matrix[guardRow][guardCol].Right = true
			guardCol++
		} else if guardFacing == 2 {
			matrix[guardRow][guardCol].Down = true
			guardRow++
		} else if guardFacing == 3 {
			matrix[guardRow][guardCol].Left = true
			guardCol--
		}

		// Check if guard has left the map again because he was moved
		if guardRow < 0 || guardRow >= len(matrix) || guardCol < 0 || guardCol >= len(matrix[0]) {
			break
		}

		// Guard completed a loop
		if guardFacing == 0 && matrix[guardRow][guardCol].Up {
			break
		} else if guardFacing == 1 && matrix[guardRow][guardCol].Right {
			break
		} else if guardFacing == 2 && matrix[guardRow][guardCol].Down {
			break
		} else if guardFacing == 3 && matrix[guardRow][guardCol].Left {
			break
		}

		// Move the guard back and turn him if there is something in the way
		if matrix[guardRow][guardCol].Blocked {
			fmt.Println(guardRow, guardCol)
			if guardFacing == 0 {
				guardRow++
			} else if guardFacing == 1 {
				guardCol--
			} else if guardFacing == 2 {
				guardRow--
			} else if guardFacing == 3 {
				guardCol++
			}
			guardFacing = (guardFacing + 1) % 4
		}
	}

	fmt.Println(visitedCount)
}
