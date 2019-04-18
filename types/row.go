// Package types defines all available word search types.
package types

// Row defines a cross word puzzle row.
type Row struct {
	Letters string
}

/* BEGIN EXPORTED METHODS */

// NewRow initializes a new word search row.
func NewRow(letters string) *Row {
	return &Row{
		Letters: letters, // Set letters
	} // Return initialized row
}

// SetLetters sets all letters in a given row.
func (row *Row) SetLetters(letters string) {
	(*row).Letters = letters // Set letters
}

/* END EXPORTED METHODS */
