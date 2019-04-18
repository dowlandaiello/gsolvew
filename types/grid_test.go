// Package types defines all available word search types.
package types

import "testing"

/* BEGIN EXPORTED METHODS TESTS */

// TestFindString tests the functionality of the FindString() method.
func TestFindString(t *testing.T) {
	grid := NewGrid(3, 3) // Initialize grid

	rows := []string{"abc", "def", "ghi"} // Init rows

	for i, row := range rows { // Iterate through rows
		grid.Rows[i].SetLetters(row) // Set letters
	}

	x, y := grid.FindString("abc") // Find string

	if x[0] != 0 || y[0] != 0 { // Check invalid coords
		t.Fatal("invalid coords") // Panic
	}

	if x[1] != 1 || y[0] != 0 { // Check invalid coords
		t.Fatal("invalid coords") // Panic
	}

	x, y = grid.FindString("ae") // Find diagonal

	if x[0] != 0 || y[0] != 0 { // Check invalid coords
		t.Fatal("invalid coords") // Panic
	}

	if x[1] != 1 || y[1] != 1 { // Check invalid coords
		t.Fatal("invalid coords") // Panic
	}
}

// TestGetCharPos tests the functionality of the GetCharPos() method.
func TestGetCharPos(t *testing.T) {
	grid := NewGrid(3, 3) // Initialize grid

	rows := []string{"abc", "def", "ghi"} // Init rows

	for i, row := range rows { // Iterate through rows
		grid.Rows[i].SetLetters(row) // Set letters
	}

	x, y := grid.GetCharPos('a', 0) // Get pos

	if x != 0 || y != 0 { // Check invalid
		t.Fatal("failed to get pos") // Panic
	}

	x, y = grid.GetCharPos('b', 0) // Get pos

	if x != 1 || y != 0 { // Check invalid
		t.Fatal("failed to get pos") // Panic
	}

	x, y = grid.GetCharPos('e', 0) // Get pos

	if x != 1 || y != 1 { // Check invalid
		t.Fatal("failed to get pos") // Panic
	}
}

/* END EXPORTED METHODS TESTS */
