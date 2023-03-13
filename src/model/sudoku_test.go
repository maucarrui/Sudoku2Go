package model

import (
	"testing"
)

// getValidSudokus returns an array of valid sudoku values.
func getValidSudokus() [][]int {
	// List of valid sudoku values.
	validSudokus := [][]int{
		// Empty sudoku.
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0},

		// Completed sudoku.
		{
			6, 3, 8, 2, 4, 7, 1, 9, 5,
			4, 2, 5, 8, 1, 9, 7, 6, 3,
			7, 9, 1, 5, 3, 6, 4, 8, 2,
			9, 5, 7, 1, 8, 3, 6, 2, 4,
			3, 8, 2, 7, 6, 4, 9, 5, 1,
			1, 6, 4, 9, 5, 2, 3, 7, 8,
			2, 4, 9, 3, 7, 5, 8, 1, 6,
			8, 7, 6, 4, 2, 1, 5, 3, 9,
			5, 1, 3, 6, 9, 8, 2, 4, 7},

		// Non-completed sudoku.
		{
			6, 0, 8, 0, 4, 0, 1, 0, 5,
			0, 2, 0, 8, 0, 9, 0, 6, 3,
			7, 0, 1, 0, 3, 0, 4, 8, 0,
			0, 5, 0, 1, 0, 3, 6, 0, 4,
			3, 0, 2, 0, 6, 4, 0, 5, 0,
			0, 6, 0, 9, 5, 0, 3, 0, 8,
			2, 0, 9, 3, 0, 5, 0, 1, 0,
			0, 7, 6, 0, 2, 0, 5, 0, 9,
			5, 1, 0, 6, 0, 8, 0, 4, 0},
	}

	return validSudokus
}

// getInvalidSudokus returns an array of invalid sudoku values.
func getInvalidSudokus() [][]int {
	// List of invalid sudoku values.
	invalidSudokus := [][]int{

		// No repetitions on the same row.
		{
			1, 1, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0},

		// No repetitions on the same column.
		{
			1, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 0, 0, 0, 0, 0, 0, 0, 0},

		// No repetitions on the same block.
		{
			1, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0},

		// No values greater than 9.
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 10, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0},

		// No values lesser than 0.
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, -1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	return invalidSudokus
}

// TestNewSudoku tests if NewSudoku returns a proper sudoku structure or a
// proper error.
func TestNewSudoku(t *testing.T) {
	err_msg_start := "Sudoku.NewSudoku: "
	err_msg := ""

	// Check that a sudoku is initialized when given a correct array of values.
	err_msg = err_msg_start
	err_msg += "valid values %v returned an error "
	err_msg += "(returned error: %v)"
	for _, values := range getValidSudokus() {
		_, err := NewSudoku(values)

		if err != nil {
			t.Errorf(err_msg, values, err)
		}
	}

	// Check that when given invalid values, an error is returned.
	err_msg = err_msg
	err_msg += "invalid values are accepted: %v"
	for _, values := range getInvalidSudokus() {
		_, err := NewSudoku(values)

		if err == nil {
			t.Errorf(err_msg, values)
		}
	}
}

// TestSetValue tests if SetValue properly sets a value in sudoku if valid, and
// if not valid it tests if it returns an error.
func TestSetValue(t *testing.T) {
	emptyEntries := getValidSudokus()[0]

	err_msg_start := "Sudoku.SetValue: "
	err_msg := ""

	sudoku, _ := NewSudoku(emptyEntries)

	// Must return error if value is below 1.
	err_msg = err_msg_start + "accepts values below 1: %d"
	for _, val := range [4]int{0, -10, -40, -100} {
		if err := sudoku.SetValue(0, 0, val); err == nil {
			t.Errorf(err_msg, val)
		}
	}

	// Must return error if value is above 9.
	err_msg = err_msg_start + "accepts values below 9: %d"
	for _, val := range [4]int{10, 100, 20, 34} {
		if err := sudoku.SetValue(0, 0, val); err == nil {
			t.Errorf(err_msg, val)
		}
	}

	// Must not return error for a value between 1 and 9.
	err_msg = err_msg_start
	err_msg += "doesn't accept values between 1-9: %d "
	err_msg += "(returned error: %v)"
	for i := 1; i <= 9; i++ {
		if err := sudoku.SetValue(0, 0, i); err != nil {
			t.Errorf(err_msg, i, err)
		}
	}

	// Must return an error if we try to write in an invalid row.
	err_msg = err_msg_start + "invalid row %d is accepted"
	for _, val := range [4]int{-1, -2, 10, 11} {
		if err := sudoku.SetValue(val, 0, 1); err == nil {
			t.Errorf(err_msg, val)
		}
	}

	// Must return an error if we try to write in an invalid column.
	err_msg = err_msg_start + "invalid column %d is accepted"
	for _, val := range [4]int{-1, -2, -10, -100} {
		if err := sudoku.SetValue(0, val, 1); err == nil {
			t.Errorf(err_msg, val)
		}
	}

	sudoku, _ = NewSudoku(emptyEntries)

	// Must not return an error if we try to write in a valid row.
	err_msg = err_msg_start
	err_msg += "returns an error in valid row %d "
	err_msg += "(returned error: %v)"
	for i := 0; i < 9; i++ {
		if err := sudoku.SetValue(i, 0, i+1); err != nil {
			t.Errorf(err_msg, i, err)
		}
	}

	sudoku, _ = NewSudoku(emptyEntries)

	// Must not return an error if we try to write in a valid column.
	err_msg = err_msg_start
	err_msg += "returns error in valid column %d "
	err_msg += "(returned error: %v)"
	for i := 0; i < 9; i++ {
		if err := sudoku.SetValue(0, i, i+1); err != nil {
			t.Errorf(err_msg, i, err)
		}
	}

	sudoku, _ = NewSudoku(emptyEntries)

	// Must return an error when writing the same value in the same row.
	err_msg = err_msg_start
	err_msg += "accepts repeated values on rows"
	sudoku.SetValue(0, 0, 1)
	if err := sudoku.SetValue(0, 1, 1); err == nil {
		t.Errorf(err_msg)
	}

	// Must return an error when writing the same value in the same column.
	err_msg = err_msg_start
	err_msg += "accepts repeated values on columns"
	if err := sudoku.SetValue(1, 0, 1); err == nil {
		t.Errorf(err_msg)
	}

	// Must return an error when writing the same value in the same block.
	err_msg = err_msg_start
	err_msg += "accepts repeated values on blocks"
	if err := sudoku.SetValue(2, 2, 1); err == nil {
		t.Errorf(err_msg)
	}
}
