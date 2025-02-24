package model

import (
	"fmt"                                     // Error formatting.
	"github.com/maucarrui/Sudoku2Go/internal" // Integer Set structure.
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
func NewSudoku(values []int) (*Sudoku, *SudokuError) {
	// If the array doesn't have 81 elements, return an error.
	if len(values) != 81 {
		return nil, NewSudokuError().
			SetErrorType(INVALID_SUDOKU).
			SetErrorMessage(fmt.Sprintf("Need 81 elements, found %d", len(values)))
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
			return nil, NewSudokuError().
				SetErrorType(INVALID_SUDOKU).
				SetStackError(err).
				SetErrorMessage("Invalid Sudoku")
		}
	}

	return s, nil
}

// SetValue sets a value in the sudoku in the given row and column. In case the
// value is greater than 9 or less than 0, or it is repeated in a row, column or
// block, return an error.
func (sudoku *Sudoku) SetValue(row, col, val int) *SudokuError {
	// Check for valid row, column, and value.
	if row < 0 || row > 8 {
		return NewSudokuError().
			SetErrorType(INVALID_ROW).
			SetErrorMessage(fmt.Sprintf("Invalid row %d", row))
	}

	if col < 0 || col > 8 {
		return NewSudokuError().
			SetErrorType(INVALID_COLUMN).
			SetErrorMessage(fmt.Sprintf("Invalid column %d", col))
	}

	if val < 1 || val > 9 {
		return NewSudokuError().
			SetErrorType(INVALID_VALUE).
			SetErrorMessage(fmt.Sprintf("Invalid value %d", val))
	}

	// Check for no repetition on row, column and block.
	if sudoku.rowVals[row].Contains(val) {
		conflictColumn := sudoku.findConflictingColumn(row, val)
		return NewSudokuError().
			SetErrorType(REPEATED_VALUE).
			SetConflictRow(row).
			SetConflictColumn(conflictColumn).
			SetErrorMessage(fmt.Sprintf("Repeated value %d", val))
	}

	if sudoku.colVals[col].Contains(val) {
		conflictRow := sudoku.findConflictingRow(col, val)
		return NewSudokuError().
			SetErrorType(REPEATED_VALUE).
			SetConflictRow(conflictRow).
			SetConflictColumn(col).
			SetErrorMessage(fmt.Sprintf("Repeated value %d", val))
	}

	block := ((row / 3) * 3) + (col / 3)
	if sudoku.blockVals[block].Contains(val) {
		conflictRow, conflictColumn := sudoku.findConflictingRowAndColumn(block, val)
		return NewSudokuError().
			SetErrorType(REPEATED_VALUE).
			SetConflictRow(conflictRow).
			SetConflictColumn(conflictColumn).
			SetErrorMessage(fmt.Sprintf("Repeated value %d", val))
	}

	// If the entry is valid, add it.
	sudoku.rowVals[row].Add(val)
	sudoku.colVals[col].Add(val)
	sudoku.blockVals[block].Add(val)
	sudoku.values[row][col] = val
	sudoku.emptyEntries--
	return nil
}

func (sudoku *Sudoku) RemoveValue(row, col int) *SudokuError {
	// Check for valid row, column.
	if row < 0 || row > 8 {
		return NewSudokuError().
			SetErrorType(INVALID_ROW).
			SetErrorMessage(fmt.Sprintf("Invalid row %d", row))
	}

	if col < 0 || col > 8 {
		return NewSudokuError().
			SetErrorType(INVALID_COLUMN).
			SetErrorMessage(fmt.Sprintf("Invalid column %d", col))
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
func (sudoku *Sudoku) GetValue(row, col int) (int, *SudokuError) {
	if row < 0 || row > 8 {
		return 0, NewSudokuError().
			SetErrorType(INVALID_ROW).
			SetErrorMessage(fmt.Sprintf("Invalid row %d", row))
	}

	if col < 0 || col > 8 {
		return 0, NewSudokuError().
			SetErrorType(INVALID_COLUMN).
			SetErrorMessage(fmt.Sprintf("Invalid column %d", col))
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

func (sudoku *Sudoku) findConflictingColumn(row, value int) int {
	for col := 0; col < 9; col++ {
		if sudoku.values[row][col] == value {
			return col
		}
	}
	return -1
}

func (sudoku *Sudoku) findConflictingRow(col, value int) int {
	for row := 0; row < 9; row++ {
		if sudoku.values[row][col] == value {
			return row
		}
	}
	return -1
}

func (sudoku *Sudoku) findConflictingRowAndColumn(block, value int) (int, int) {
	row := (block / 3) * 3
	col := (block % 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku.values[row+i][col+j] == value {
				return row + i, col + j
			}
		}
	}
	return -1, -1
}
