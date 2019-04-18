// Package types defines all available word search types.
package types

import (
	"math"
)

// Grid outlines the default word search grid.
type Grid struct {
	Rows []*Row // Rows
}

/* BEGIN EXPORTED METHODS */

// NewGrid initializes a new grid with a given width and height.
func NewGrid(width, height uint64) *Grid {
	grid := &Grid{
		Rows: []*Row{}, // Set rows
	} // Return init grid

	for i := uint64(0); i < height; i++ { // Make rows
		grid.Rows = append(grid.Rows, NewRow("")) // Initialize row
	}

	return grid // Return grid
}

// FindString finds the position of all requested chars in the given grid.
func (grid *Grid) FindString(s string) ([]uint64, []uint64) {
	var xCoordinates []uint64 // Init x coords buffer
	var yCoordinates []uint64 // Init y coords buffer

	n := 0 // Init iterator

findCoords:
	for i, char := range s { // Iterate through chars
		x, y := grid.GetCharPos(char, n) // Get position

		if x == -1 || y == -1 { // Check invalid coords
			return []uint64{}, []uint64{} // Return null
		}

		if i != 0 { // Check is not first char
			if math.Abs(float64(xCoordinates[i-1])-float64(x)) > 1 { // Check invalid selection
				xCoordinates = xCoordinates[:0] // Reset slice
				yCoordinates = yCoordinates[:0] // Reset slice

				n++ // Increment

				goto findCoords // Reset iteration
			} else if math.Abs(float64(yCoordinates[i-1])-float64(y)) > 1 { // Check invalid selection
				xCoordinates = xCoordinates[:0] // Reset slice
				yCoordinates = yCoordinates[:0] // Reset slice

				n++ // Increment

				goto findCoords // Reset iteration
			}
		}

		xCoordinates = append(xCoordinates, uint64(x)) // Append x
		yCoordinates = append(yCoordinates, uint64(y)) // Append y
	}

	return xCoordinates, yCoordinates // Return coordinates
}

// GetCharPos gets the position of a char in the given grid.
func (grid *Grid) GetCharPos(c rune, n int) (int64, int64) {
	i := 0 // Init iterator

	for y, row := range grid.Rows { // Iterate through rows
		for x, char := range row.Letters { // Iterate through row letters
			if char == c { // Check matching char
				if i != n { // Check invalid match
					i++ // Increment iterator

					continue // Continue
				}

				return int64(x), int64(y) // Return coords
			}
		}
	}

	return 0, 0 // Invalid
}

// GetRow gets a reference to the row at iterator i.
func (grid *Grid) GetRow(i uint64) *Row {
	return grid.Rows[i] // Return row
}

/* END EXPORTED METHODS */

/* BEGIN INTERNAL METHODS */

func makeString(len uint64) string {
	var str string // Init string buffer

	for i := uint64(0); i < len; i++ { // Make string
		str = str + "0" // Append empty string
	}

	return str // Retrun string
}

/* END INTERNAL METHODS */
