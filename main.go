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
	if len(os.Args) == 1 { // Check invalid params
		panic("Make sure to pass in word search width and height!") // Panic
	}

	gridWidth, err := strconv.Atoi(os.Args[1]) // Parse width

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	gridHeight, err := strconv.Atoi(os.Args[2]) // Parse height

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	grid := types.NewGrid(uint64(gridWidth), uint64(gridHeight)) // Initialize grid

	reader := bufio.NewReader(os.Stdin) // Initialize reader

	for i := 0; i < gridHeight; i++ { // Do until made all rows
		fmt.Printf("Row %d: ", i) // Log input row

		input, err := reader.ReadString('\n') // Read up to \n

		if err != nil { // Check for errors
			panic(err) // Panic
		}

		input = strings.Replace(input, "\n", "", -1) // Remove \n

		grid.Rows[i].Letters = input // Set letters
	}

	for {
		fmt.Print("Word to find: ") // Print prompt

		wordToFind, err := reader.ReadString('\n') // Read up to \n

		if err != nil { // Check for errors
			panic(err) // Panic
		}

		wordToFind = strings.Replace(wordToFind, "\n", "", -1) // Remove \n

		xCoords, yCoords := grid.FindString(wordToFind) // Find word

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

// indexOf fetches the index of a value in a uint64 slice.
func indexOf(s []uint64, v uint64) int {
	for i, val := range s { // Iterate through values
		if val == v { // Check match
			return i // Valid
		}
	}

	return -1 // Invalid
}
