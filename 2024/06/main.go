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

func simulateRoute(matrix [][]Position, guardRow, guardCol, guardFacing int) (int, bool) {
	visitedCount := 0
	for {
		// Guard completed a loop
		if guardFacing == 0 && matrix[guardRow][guardCol].Up {
			return visitedCount, true
		} else if guardFacing == 1 && matrix[guardRow][guardCol].Right {
			return visitedCount, true
		} else if guardFacing == 2 && matrix[guardRow][guardCol].Down {
			return visitedCount, true
		} else if guardFacing == 3 && matrix[guardRow][guardCol].Left {
			return visitedCount, true
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

		// Check if guard has left the map because he was moved
		if guardRow < 0 || guardRow >= len(matrix) || guardCol < 0 || guardCol >= len(matrix[0]) {
			return visitedCount, false
		}

		// Move the guard back and turn him if there is something in the way
		// Don't reset the positions guard facing history. a turn is still
		// a movement
		if matrix[guardRow][guardCol].Blocked {
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
}

func deepCopyMatrix(matrix [][]Position) [][]Position {
	newMatrix := make([][]Position, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]Position, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
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
			pos := Position{Blocked: false, Left: false, Right: false, Up: false, Down: false}
			if char == '^' {
				guardRow = rowCount
				guardCol = col
				guardFacing = 0
			} else if char == '>' {
				guardRow = rowCount
				guardCol = col
				guardFacing = 1
			} else if char == 'v' {
				guardRow = rowCount
				guardCol = col
				guardFacing = 2
			} else if char == '<' {
				guardRow = rowCount
				guardCol = col
				guardFacing = 3
			} else if char == '#' {
				pos.Blocked = true
			}
			row = append(row, pos)
		}
		matrix = append(matrix, row)
		rowCount++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	visited, _ := simulateRoute(deepCopyMatrix(matrix), guardRow, guardCol, guardFacing)

	loops := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if !(row == guardRow && col == guardCol) && !matrix[row][col].Blocked {
				matrix[row][col].Blocked = true
				if _, loop := simulateRoute(deepCopyMatrix(matrix), guardRow, guardCol, guardFacing); loop {
					loops++
				}
				matrix[row][col].Blocked = false
			}
		}
	}

	fmt.Println(visited)
	fmt.Println(loops)
}
