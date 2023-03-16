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

// TestSetValueInvalidEntries tests if SetValue properly returns an error when
// given an invalid value, row, or column.
func TestSetValueInvalidEntries(t *testing.T) {
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

	// Must return an error if we try to write in an invalid row.
	err_msg = err_msg_start + "invalid row is accepted: %d"
	for _, val := range [4]int{-1, -2, 10, 11} {
		if err := sudoku.SetValue(val, 0, 1); err == nil {
			t.Errorf(err_msg, val)
		}
	}

	// Must return an error if we try to write in an invalid column.
	err_msg = err_msg_start + "invalid column is accepted: %d"
	for _, val := range [4]int{-1, -2, -10, -100} {
		if err := sudoku.SetValue(0, val, 1); err == nil {
			t.Errorf(err_msg, val)
		}
	}
}

// TestSetValueValidEntries tests if SetValue doesn't return an error when given
// valid entries as values, rows, and columns.
func TestSetValueValidEntries(t *testing.T) {
	emptyEntries := getValidSudokus()[0]
	err_msg_start := "Sudoku.SetValue: "
	err_msg := ""

	sudoku, _ := NewSudoku(emptyEntries)

	// Must not return error for a value between 1 and 9.
	err_msg = err_msg_start
	err_msg += "doesn't accept values between 1-9: %d "
	err_msg += "(returned error: %v)"
	for i := 1; i <= 9; i++ {
		if err := sudoku.SetValue(0, 0, i); err != nil {
			t.Errorf(err_msg, i, err)
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

// TestRemoveValue tests if RemoveValue properly removes an entry in the sudoku.
func TestRemoveValue(t *testing.T) {
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
