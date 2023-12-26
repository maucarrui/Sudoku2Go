package model

import (
	"fmt"                                     // Error formatting.
	"github.com/maucarrui/Sudoku2Go/internal" // Integer Set structure.
	"strconv"                                 // String Conversions.
)

// Alias for Integer set structure.
type IntSet = internal.IntSet

type Sudoku struct {
	// All the sudoku values.
	values [9][9]int

	// Row map to keep track of the values contained in each row.
	rowVals map[int]*IntSet

	// Column map to keep track of the values contained in each column.
	colVals map[int]*IntSet

	// Block map to keep track of the values contained in each block.
	blockVals map[int]*IntSet

	// Amount of empty entries in the sudoku.
	emptyEntries int
}

// NewSudoku initializes a Sudoku with the given values. The given values are an
// array of integers, such that (i / 10) and (i % 10) determine the row and
// column where the values belong in the sudoku, respectively, where i is the
// index of the value in the array. The value 0 is considered as an empty cell,
// and thus it should be ignored. The array has to contain exactly 81 elements.
func NewSudoku(values []int) (*Sudoku, error) {
	// If the array doesn't have 81 elements, return an error.
	if len(values) != 81 {
		err_str := "can't initialize sudoku: need 81 elements, found %d"
		return nil, fmt.Errorf(err_str, len(values))
	}

	// Initialize an empty sudoku.
	s := &Sudoku{
		rowVals:      make(map[int]*IntSet),
		colVals:      make(map[int]*IntSet),
		blockVals:    make(map[int]*IntSet),
		emptyEntries: 81,
	}

	// Initialize an empty set for each row, column, and block map.
	for i := 0; i < 10; i++ {
		s.rowVals[i] = internal.NewIntSet()
		s.colVals[i] = internal.NewIntSet()
		s.blockVals[i] = internal.NewIntSet()
	}

	// Insert the values into the sudoku.
	for i, val := range values {
		// If the value is 0, ignore it.
		if val == 0 {
			continue
		}

		// Determine the row and column of the value.
		row := (i / 9)
		col := (i % 9)

		// Insert the value in the sudoku and check for errors.
		err := s.SetValue(row, col, val)
		if err != nil {
			err_str := "can't initialize sudoku: %w"
			return nil, fmt.Errorf(err_str, err)
		}
	}

	return s, nil
}

// SetValue sets a value in the sudoku in the given row and column. In case the
// value is greater than 9 or less than 0, or it is repeated in a row, column or
// block, return an error.
func (sudoku *Sudoku) SetValue(row, col, val int) error {
	// Check for valid row, column, and value.
	if row < 0 || row > 8 {
		return fmt.Errorf("invalid row %d", row)
	}

	if col < 0 || col > 8 {
		return fmt.Errorf("invalid column %d", col)
	}

	if val < 1 || val > 9 {
		return fmt.Errorf("invalid value %d", val)
	}

	// Check for no repetition on row, column and block.
	if sudoku.rowVals[row].Contains(val) {
		return fmt.Errorf("repeated value %d in row %d", val, row)
	}

	if sudoku.colVals[col].Contains(val) {
		return fmt.Errorf("repeated value %d in column %d", val, col)
	}

	block := ((row / 3) * 3) + (col / 3)
	if sudoku.blockVals[block].Contains(val) {
		return fmt.Errorf("repeated value %d in block %d", val, block)
	}

	// If the entry is valid, add it.
	sudoku.rowVals[row].Add(val)
	sudoku.colVals[col].Add(val)
	sudoku.blockVals[block].Add(val)
	sudoku.values[row][col] = val
	sudoku.emptyEntries--
	return nil
}

func (sudoku *Sudoku) RemoveValue(row, col int) error {
	// Check for valid row, column.
	if row < 0 || row > 8 {
		return fmt.Errorf("invalid row %d", row)
	}

	if col < 0 || col > 8 {
		return fmt.Errorf("invalid column %d", col)
	}

	// If the entry was non-empty, set it to zero and increase the amount of empty
	// entries.
	if sudoku.values[row][col] != 0 {
		block := ((row / 3) * 3) + (col / 3)
		val := sudoku.values[row][col]

		sudoku.rowVals[row].Remove(val)
		sudoku.colVals[col].Remove(val)
		sudoku.blockVals[block].Remove(val)

		sudoku.values[row][col] = 0
		sudoku.emptyEntries++
	}

	return nil
}

// GetValue returns the value of the sudoku found in the given row and column.
func (sudoku *Sudoku) GetValue(row, col int) (int, error) {
	if row < 0 || row > 8 {
		return 0, fmt.Errorf("invalid row %d", row)
	}

	if col < 0 || col > 8 {
		return 0, fmt.Errorf("invalid column %d", col)
	}

	return sudoku.values[row][col], nil
}

// GetValues returns the matrix representation of the sudoku, that is, a 9x9
// matrix.
func (sudoku *Sudoku) GetValues() [9][9]int {
	return sudoku.values
}

// IsComplete returns true if the sudoku all its entries are non-empty.
func (sudoku *Sudoku) IsComplete() bool {
	return sudoku.emptyEntries == 0
}

// Prints a part of a grid with (length) squares in a row.
// Part 1: Upper Grid.
// Part 2: Middle Grid.
// Part 3: Bottom Grid.
func printGridPart(part, length int, delim bool) string {
	var gridString string

	switch part {
	case 1:
		gridString = "╔"
	case 2:
		if delim {
			gridString = "╠"
		} else {
			gridString = "├"
		}
	case 3:
		gridString = "╚"
	}

	for i := 0; i < length; i++ {
		gridString += "───"

		if i < (length - 1) {
			if delim && (i+1)%3 == 0 {
				switch part {
				case 1:
					gridString += "╦"
				case 2:
					gridString += "╬"
				case 3:
					gridString += "╩"
				}
			} else {
				switch part {
				case 1:
					gridString += "┬"
				case 2:
					gridString += "┼"
				case 3:
					gridString += "┴"
				}
			}
		} else {
			switch part {
			case 1:
				gridString += "╗"
			case 2:
				if delim {
					gridString += "╣"
				} else {
					gridString += "┤"
				}
			case 3:
				gridString += "╝"
			}
		}
	}

	return gridString
}

// Returns the given sudoku in String format.
func (sudoku *Sudoku) ToString() string {
	var delim bool
	sudokuString := printGridPart(1, 9, true) + "\n"

	for i, row := range sudoku.values {

		// Print the current row values.
		for j := 0; j < len(row); j++ {

			if (i+1)%3 == 0 {
				delim = true
			} else {
				delim = false
			}

			sudokuString += "│ "
			sudokuString += strconv.Itoa(row[j]) + " "

			if j == len(row)-1 {
				sudokuString += "│\n"
			}
		}

		// Depending on the row, print the corresponding grid part.
		if i < 8 {
			sudokuString += printGridPart(2, 9, delim) + "\n"
		} else {
			sudokuString += printGridPart(3, 9, delim) + "\n"
		}
	}

	return sudokuString
}
