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

		{
			6, 3, 8, 2, 4, 7, 1, 9, 5,
			4, 2, 5, 8, 1, 9, 7, 6, 3,
			7, 9, 1, 5, 3, 6, 4, 8, 2,
			9, 5, 7, 1, 8, 3, 6, 2, 4,
			3, 8, 2, 7, 6, 4, 9, 5, 1,
			1, 6, 4, 9, 5, 2, 3, 7, 8,
			2, 4, 9, 3, 7, 5, 8, 0, 6,
			8, 7, 6, 4, 2, 1, 5, 3, 9,
			5, 1, 3, 6, 9, 8, 2, 4, 7},
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

		// Less than 81 elements.
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0},

		// More than 81 elements.
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	return invalidSudokus
}

// getCompleteSudokus returns an array of completed sudokus.
func getCompleteSudokus() [][]int {
	completeSudokus := [][]int{
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
	}

	return completeSudokus
}

// TestNewSudokuValidValues tests if NewSudoku returns a proper sudoku structure
// for valid values.
func TestNewSudokuValidValues(t *testing.T) {
	err_msg := "Sudoku.NewSudoku: "
	err_msg += "valid values returned an error: %v "
	err_msg += "(returned error: %v)"

	// Check that a sudoku is initialized when given a correct array of values.
	for _, values := range getValidSudokus() {
		_, err := NewSudoku(values)

		if err != nil {
			t.Errorf(err_msg, values, err)
		}
	}
}

// TestNewSudokuInvalidValues tests if NewSudoku returns an error when given
// invalid values.
func TestNewSudokuInvalidValues(t *testing.T) {
	// Check that when given invalid values, an error is returned.
	err_msg := "Sudoku.NewSudoku: "
	err_msg += "invalid values are accepted: %v"
	for _, values := range getInvalidSudokus() {
		_, err := NewSudoku(values)

		if err == nil {
			t.Errorf(err_msg, values)
		}
	}
}

// TestNewSudokuConsistency tests if NewSudoku returns a sudoku which contains
// the values that were passed.
func TestNewSudokuConsistency(t *testing.T) {
	err_msg := "Sudoku.NewSudoku: "

	for _, values := range getValidSudokus() {
		sudoku, err := NewSudoku(values)

		// The initializing the sudoku must not return an error.
		if err != nil {
			err_msg += "valid values return an unexpected error: %v"
			t.Errorf(err_msg, err)
			break
		}

		for i, val := range values {
			row := (i / 9)
			col := (i % 9)

			obt_val, err := sudoku.GetValue(row, col)

			// Check for unexpected errors.
			if err != nil {
				err_msg += "valid values return an unexpected error: %v"
				t.Errorf(err_msg, err)
				break
			}

			// Check that the obtained value is correct.
			if val != obt_val {
				err_msg += "expected value %v but got %v"
				t.Errorf(err_msg, val, obt_val)
			}
		}
	}
}

// TestGetValueValidRows tests if GetValue doesn't return an error when given
// valid rows.
func TestGetValueValidRows(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	sudoku, _ := NewSudoku(emptyEntries)

	err_msg := "Sudoku.GetValue: "
	err_msg += "unexpected error when getting value at row %d: %v"

	for row := 0; row < 9; row++ {
		if _, err := sudoku.GetValue(row, 0); err != nil {
			t.Errorf(err_msg, row, err)
		}
	}
}

// TestGetValueInvalidRows tests if GetValue properly returns an error when
// given invalid rows.
func TestGetValueInvalidRows(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	sudoku, _ := NewSudoku(emptyEntries)

	err_msg_start := "Sudoku.GetValue: "
	err_msg := ""

	// Must return an error if row is below 0.
	err_msg = err_msg_start + "accepts row below 1: %d"
	for _, row := range []int{-1, -10, -100} {
		if _, err := sudoku.GetValue(row, 0); err == nil {
			t.Errorf(err_msg, row)
		}
	}

	// Must return an error if row is above 8.
	err_msg = err_msg_start + "accepts row above 8: %d"
	for _, row := range []int{9, 10, 100, 20} {
		if _, err := sudoku.GetValue(row, 0); err == nil {
			t.Errorf(err_msg, row)
		}
	}
}

// TestGetValueValidColumns tests if GetValue doesn't return an error when given
// valid columns.
func TestGetValueValidColumns(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	sudoku, _ := NewSudoku(emptyEntries)

	err_msg := "Sudoku.GetValue: "
	err_msg += "unexpected error when getting value at column %d: %v"

	for col := 0; col < 9; col++ {
		if _, err := sudoku.GetValue(0, col); err != nil {
			t.Errorf(err_msg, col, err)
		}
	}
}

// TestGetValueInvalidColumns tests if GetValue properly returns an error when
// given invalid columns.
func TestGetValueInvalidColumns(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	sudoku, _ := NewSudoku(emptyEntries)

	err_msg_start := "Sudoku.GetValue: "
	err_msg := ""

	// Must return an error if column is below 0.
	err_msg = err_msg_start + "accepts column below 1: %d"
	for _, col := range []int{-1, -10, -100} {
		if _, err := sudoku.GetValue(0, col); err == nil {
			t.Errorf(err_msg, col)
		}
	}

	// Must return an error if row is above 8.
	err_msg = err_msg_start + "accepts column above 8: %d"
	for _, col := range []int{9, 10, 100, 20} {
		if _, err := sudoku.GetValue(0, col); err == nil {
			t.Errorf(err_msg, col)
		}
	}
}

// TestSetValueInvalidValues tests if SetValue properly returns an error when
// given invalid values.
func TestSetValueInvalidValues(t *testing.T) {
	emptyEntries := getValidSudokus()[0]

	err_msg := "Sudoku.SetValue: "

	sudoku, _ := NewSudoku(emptyEntries)

	// Must return an error if value is below 1.
	err_msg += "accepts values below 1: %d"
	for _, val := range []int{0, -10, -40, -100} {
		if err := sudoku.SetValue(0, 0, val); err == nil {
			t.Errorf(err_msg, val)
		}
	}

	// Must return an error if value is above 9.
	err_msg += "accepts values above 9: %d"
	for _, val := range []int{10, 100, 20, 34} {
		if err := sudoku.SetValue(0, 0, val); err == nil {
			t.Errorf(err_msg, val)
		}
	}
}

// TestSetValueInvalidRows tests if SetValue properly returns an error when
// given invalid rows.
func TestSetValueInvalidRows(t *testing.T) {
	emptyEntries := getValidSudokus()[0]

	err_msg := "Sudoku.SetValue: "

	sudoku, _ := NewSudoku(emptyEntries)

	// Must return an error if we try to write in an invalid row.
	err_msg += "invalid row is accepted: %d"
	for _, val := range []int{-1, -2, 10, 11} {
		if err := sudoku.SetValue(val, 0, 1); err == nil {
			t.Errorf(err_msg, val)
		}
	}
}

// TestSetValueInvalidColumns tests if SetValue properly returns an error when
// given invalid columns.
func TestSetValueInvalidColumns(t *testing.T) {
	emptyEntries := getValidSudokus()[0]

	err_msg := "Sudoku.SetValue: "

	sudoku, _ := NewSudoku(emptyEntries)

	// Must return an error if we try to write in an invalid column.
	err_msg += "invalid column is accepted: %d"
	for _, val := range [4]int{-1, -2, -10, -100} {
		if err := sudoku.SetValue(0, val, 1); err == nil {
			t.Errorf(err_msg, val)
		}
	}
}

// TestSetValueNoRepetitionRows tests if SetValue checks properly for no
// repetition on the rows.
func TestSetValueNoRepetitionRows(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	err_msg_start := "Sudoku.SetValue: "
	err_msg := ""

	// Check for no repetitions on rows.
	err_msg = err_msg_start
	err_msg += "accepts repeated values on rows"
	for i := 0; i < 10; i++ {
		sudoku, _ := NewSudoku(emptyEntries)
		sudoku.SetValue(i, 0, 1)

		for j := 1; j < 10; j++ {
			if err := sudoku.SetValue(i, j, 1); err == nil {
				t.Errorf(err_msg)
			}
		}
	}
}

// TestSetValueNoRepetitionColumns tests if SetValue checks properly for no
// repetition on the columns.
func TestSetValueNoRepetitionColumns(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	err_msg_start := "Sudoku.SetValue: "
	err_msg := ""

	// Check for no repetition on columns.
	err_msg = err_msg_start
	err_msg += "accepts repeated values on columns"
	for i := 0; i < 10; i++ {
		sudoku, _ := NewSudoku(emptyEntries)
		sudoku.SetValue(0, i, 1)

		for j := 1; j < 10; j++ {
			if err := sudoku.SetValue(j, i, 1); err == nil {
				t.Errorf(err_msg)
			}
		}
	}
}

// TestSetValueNoRepetitionBlocks tests if SetValue checks properly for no
// repetition on the blocks.
func TestSetValueNoRepetitionBlocks(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	err_msg_start := "Sudoku.SetValue: "
	err_msg := ""

	// Check for no repetition on blocks.
	err_msg = err_msg_start
	err_msg += "accepts repeated values on blocks"
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sudoku, _ := NewSudoku(emptyEntries)
			sudoku.SetValue(i*3, j*3, 1)

			for k := 1; k < 3; k++ {
				for l := 1; l < 3; l++ {
					err := sudoku.SetValue((i*3 + k), (j*3 + l), 1)
					if err == nil {
						t.Errorf(err_msg)
					}
				}
			}
		}
	}
}

// TestSetValueConsistency tests if SetValue properly sets a value in the
// sudoku.
func TestSetValueConsistency(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	err_msg := "Sudoku.SetValue: "
	sudoku, err := NewSudoku(emptyEntries)

	// Check for unexpected error when initializing.
	if err != nil {
		err_msg += "unexpected error when initializing sudoku: %v"
		t.Errorf(err_msg, err)
	}

	for i := 0; i < 9; i++ {
		// Check for unexpected errors.
		if err := sudoku.SetValue(i, i, i+1); err != nil {
			err_msg += "unexpected error when setting value: %v"
			t.Errorf(err_msg, err)
			break
		}

		// Check that the placed value matches the expected.
		obt_val, err := sudoku.GetValue(i, i)
		if err != nil {
			err_msg += "unexpected error when obtaining value: %v"
			t.Errorf(err_msg, err)
			break
		}

		if obt_val != (i + 1) {
			err_msg += "expected value %d but got %d"
			t.Errorf(err_msg, (i + 1), obt_val)
			break
		}
	}
}

// Tests that SetValue properly works when first trying to place an invalid
// value, and then placing a valid value in the same row.
func TestSetValueInvalidPlacementThenValidPlacement(t *testing.T) {
	nonCompleteSudoku := getValidSudokus()[1]
	err_msg := "Sudoku.SetValue: "
	sudoku, err := NewSudoku(nonCompleteSudoku)

	// Check for unexpected error when initializing.
	if err != nil {
		err_msg += "unexpected error when initializing sudoku: %v"
		t.Errorf(err_msg, err)
	}

	// Try to place a value in an invalid position.
	if err := sudoku.SetValue(1, 2, 7); err == nil {
		err_msg += "expected an error when placing an invalid value"
		t.Errorf(err_msg)
		return
	}

	// Then place the same value but in a valid position.
	if err := sudoku.SetValue(5, 2, 7); err != nil {
		err_msg += "unexpected error when setting valid value: %v"
		t.Errorf(err_msg, err)
	}
}

// TestRemoveValueInvalidRow tests if RemoveValue properly returns an
// error when given an invalid row as an arguments.
func TestRemoveValueInvalidRow(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	sudoku, _ := NewSudoku(emptyEntries)
	err_msg := "Sudoku.RemoveValue: "
	err_msg += "invalid row returned no error: %d"

	for _, row := range [6]int{-1, -2, 9, 10, 100, -100} {
		if err := sudoku.RemoveValue(row, 0); err == nil {
			t.Errorf(err_msg, row)
		}
	}
}

// TestRemoveValueInvalidColumn tests if RemoveValue properly returns an
// error when given an invalid column as an arguments.
func TestRemoveValueInvalidColumn(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	sudoku, _ := NewSudoku(emptyEntries)
	err_msg := "Sudoku.RemoveValue: "
	err_msg += "invalid column returned no error: %d"

	for _, col := range [6]int{-1, -2, 9, 10, 100, -100} {
		if err := sudoku.RemoveValue(0, col); err == nil {
			t.Errorf(err_msg, col)
		}
	}
}

// TestRemoveValueConsistency tests if RemoveValue properly removes an entry in
// the sudoku.
func TestRemoveValueConsistency(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	err_msg_start := "Sudoku.RemoveValue: "
	err_msg := ""

	sudoku, _ := NewSudoku(emptyEntries)
	val := 0

	// For every entry, place every possible value, and then remove it.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 1; k < 10; k++ {
				err_msg = err_msg_start
				err_msg += "unexpected error when setting value: %v"
				if err := sudoku.SetValue(i, j, k); err != nil {
					t.Errorf(err_msg, err)
				}

				err_msg = err_msg_start
				err_msg += "expected value to be %d but got %d"
				if val, _ = sudoku.GetValue(i, j); val != k {
					t.Errorf(err_msg, k, val)
				}

				err_msg = err_msg_start
				err_msg += "unexpected error when removing value: %v"
				if err := sudoku.RemoveValue(i, j); err != nil {
					t.Errorf(err_msg, err)
				}

				err_msg = err_msg_start
				err_msg += "expected value to be %d but got %d"
				if val, _ = sudoku.GetValue(i, j); val != 0 {
					t.Errorf(err_msg, 0, val)
				}
			}
		}
	}
}

// TestIsCompleteTrueCase tests if IsComplete returns true when given a complete
// sudoku.
func TestIsCompleteTrueCase(t *testing.T) {
	err_msg_start := "Sudoku.IsComplete: "
	err_msg := ""

	for _, completeSudoku := range getCompleteSudokus() {
		sudoku, err := NewSudoku(completeSudoku)

		// The values must not be considered invalid.
		if err != nil {
			err_msg = err_msg_start
			err_msg += "valid values returned an error: %v"
			t.Errorf(err_msg, err)
		}

		// Check that the sudoku is complete.
		if !sudoku.IsComplete() {
			err_msg = err_msg_start
			err_msg += "the following sudoku is not considered complete: %v"
			t.Errorf(err_msg, completeSudoku)
		}
	}
}

// TestIsCompleteFalseCase tests if IsComplete returns false when given an
// incomplete sudoku.
func TestIsCompleteFalseCase(t *testing.T) {
	err_msg_start := "Sudoku.IsComplete: "
	err_msg := ""

	for _, values := range getValidSudokus() {
		sudoku, err := NewSudoku(values)

		// The values must not be considered invalid.
		if err != nil {
			err_msg = err_msg_start
			err_msg += "valid values returned an error: %v"
			t.Errorf(err_msg, err)
		}

		// Check that the sudoku is complete.
		if sudoku.IsComplete() {
			err_msg = err_msg_start
			err_msg += "the following sudoku is considered complete: %v"
			t.Errorf(err_msg, values)
		}
	}
}

// TestIsCompleteMultipleAdditions tests if IsComplete returns false even when
// adding the same element 81 times.
func TestIsCompleteMultipleAdditions(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	err_msg := "Sudoku.IsComplete: "
	sudoku, err := NewSudoku(emptyEntries)

	// Check for unexpected errors.
	if err != nil {
		err_msg += "unexpected error when initializing sudoku: %v"
		t.Errorf(err_msg, err)
	}

	// Add 81 times the same element to the same entry, the sudoku should not be
	// considered complete.
	for i := 0; i < 81; i++ {
		sudoku.SetValue(0, 0, 1)
	}

	if sudoku.IsComplete() {
		err_msg += "incomplete sudoku is considered complete"
		err_msg += "after adding the same element 81 times"
		t.Errorf(err_msg)
	}
}

// TestIsCompleteMultipleRemovals tests if IsComplete returns true even when
// removing the same element 81 times, and then simply adding it.
func TestIsCompleteMultipleRemovals(t *testing.T) {
	completeEntries := getCompleteSudokus()[0]
	err_msg := "Sudoku.IsComplete: "
	sudoku, err := NewSudoku(completeEntries)

	// Check for unexpected errors.
	if err != nil {
		err_msg += "unexpected error when initializing sudoku: %v"
		t.Errorf(err_msg, err)
	}

	// Remove 81 times the same entry, the sudoku should not be considered
	// complete.
	for i := 0; i < 81; i++ {
		sudoku.RemoveValue(0, 0)
	}

	if sudoku.IsComplete() {
		err_msg += "incomplete sudoku is considered complete"
		err_msg += "after removing the same element 81 times"
		t.Errorf(err_msg)
	}

	// Add the removed entry and the sudoku should be considered complete.
	if err := sudoku.SetValue(0, 0, completeEntries[0]); err != nil {
		err_msg += "unexpected error when adding element %d at (%d, %d) : %v"
		t.Errorf(err_msg, completeEntries[0], 0, 0, err)
	}

	// The sudoku should be complete.
	if !sudoku.IsComplete() {
		err_msg += "removed the same element 81 times, "
		err_msg += "and after adding it again the sudoku "
		err_msg += "should be considered complete, but it isn't"
		t.Errorf(err_msg)
	}
}

// TestGetValuesConsistency tests if GetValues returns the values of the sudoku
// in the proper order.
func TestGetValuesConsistency(t *testing.T) {
	complete := getCompleteSudokus()[0]
	err_msg := "Sudoku.GetValues: "

	sudoku, err := NewSudoku(complete)

	// Check for an unexpected error.
	if err != nil {
		err_msg += "unexpected error when initializing sudoku: %v"
		t.Errorf(err_msg, err)
	}

	// Check that the obtained values match the expected.
	obt_values := sudoku.GetValues()
	for i, val := range complete {
		row := (i / 9)
		col := (i % 9)

		obt_val := obt_values[row][col]

		if obt_val != val {
			err_msg += "expected value %d at (%d, %d) but got %d"
			t.Errorf(err_msg, val, row, col, obt_val)
		}
	}
}
