package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dowlandaiello/word_search/types"
	"github.com/gookit/color"
)

// main is the main word search solver function.
func main() {
	var grid *types.Grid // Declare grid buffer

	reader := bufio.NewReader(os.Stdin) // Initialize reader

	if len(os.Args) == 1 { // Check invalid params
		panic("Make sure to pass in word search width and height!") // Panic
	} else if len(os.Args) == 2 { // Check import grid
		grid = importGrid(os.Args[1]) // Import grid
	} else {
		gridWidth, err := strconv.Atoi(os.Args[1]) // Parse width

		if err != nil { // Check for errors
			panic(err) // Panic
		}

		gridHeight, err := strconv.Atoi(os.Args[2]) // Parse height

		if err != nil { // Check for errors
			panic(err) // Panic
		}

		grid = types.NewGrid(uint64(gridWidth), uint64(gridHeight)) // Initialize grid

		for i := 0; i < gridHeight; i++ { // Do until made all rows
			fmt.Printf("Row %d: ", i) // Log input row

			input, err := reader.ReadString('\n') // Read up to \n

			if err != nil { // Check for errors
				panic(err) // Panic
			}

			input = strings.Replace(input, "\n", "", -1) // Remove \n

			grid.Rows[i].Letters = input // Set letters
		}
	}

	for {
		fmt.Print("Word to find: ") // Print prompt

		wordToFind, err := reader.ReadString('\n') // Read up to \n

		if err != nil { // Check for errors
			panic(err) // Panic
		}

		wordToFind = strings.Replace(wordToFind, "\n", "", -1) // Remove \n

		xCoords, yCoords := grid.FindString(wordToFind) // Find word

		if len(xCoords) == 0 || len(yCoords) == 0 { // Check couldn't find a match
			color.Red.Println("Couldn't find a match!") // Log err

			continue // Continue
		}

		for y, row := range grid.Rows { // Iterate through rows
			var rowStr string // Init row buffer

			for x, char := range row.Letters { // Iterate through characters
				if index := indexOf(xCoords, uint64(x)); index != -1 && yCoords[index] == uint64(y) { // Check matching coordinates
					rowStr = rowStr + color.Yellow.Sprintf("%c", char) // Append char
				} else {
					rowStr = rowStr + string(char) // Append char
				}
			}

			fmt.Println(rowStr) // Log row
		}
	}
}

// importGrid imports a grid.
func importGrid(fileName string) *types.Grid {
	file, err := os.Open(fileName) // Open file

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	defer file.Close() // Close file

	scanner := bufio.NewScanner(file) // Initialize scanner

	var grid types.Grid // Initialize grid buffer

	for scanner.Scan() { // Read entire file
		line := scanner.Text() // Get line

		grid.Rows = append(grid.Rows, types.NewRow(line)) // Set row
	}

	return &grid // Return read grid
}

// indexOf fetches the index of a value in a uint64 slice.
func indexOf(s []uint64, v uint64) int {
	for i, val := range s { // Iterate through values
		if val == v { // Check match
			return i // Valid
		}
	}

	return -1 // Invalid
}
