package main

import (
	"testing"
)

func TestSetInitialValue(t *testing.T) {
	var sudoku Sudoku

	// Must return error if value is below 1.
	for _, val := range [4]int{0, -10, -40, -100} {
		if err := sudoku.SetInitialValue(0, 0, val); err == nil {
			t.Errorf("Sudoku: Accepts value below 1: %d", val)
		}
	}

	// Must return error if value is above 9.
	for _, val := range [4]int{10, 100, 20, 34} {
		if err := sudoku.SetInitialValue(0, 0, val); err == nil {
			t.Errorf("Sudoku: Accepts value above 9: %d", val)
		}
	}

	// Must not return error for a value between 1 and 9.
	for i := 1; i <= 9; i++ {
		if err := sudoku.SetInitialValue(0, 0, i); err != nil {
			t.Errorf("Sudoku: Doesn't accept values between 1-9: %d", i)
		}
	}

	// Must return an error if we try to write in an invalid row.
	for _, val := range [4]int{-1, -2, 10, 11} {
		if err := sudoku.SetInitialValue(val, 0, 1); err == nil {
			t.Errorf("Sudoku: Invalid row %d", val)
		}
	}

	// Must return an error if we try to write in an invalid column.
	for _, val := range [4]int{-1, -2, -10, -100} {
		if err := sudoku.SetInitialValue(0, val, 1); err == nil {
			t.Errorf("Sudoku: Invalid column %d", val)
		}
	}

	// Must not return an error if we try to write in a valid row.
	for i := 0; i < 9; i++ {
		if err := sudoku.SetInitialValue(i, 0, 1); err != nil {
			t.Errorf("Sudoku: Returns error in valid row %d", i)
		}
	}

	// Must not return an error if we try to write in a valid column.
	for i := 0; i < 9; i++ {
		if err := sudoku.SetInitialValue(0, i, 1); err != nil {
			t.Errorf("Sudoku: Returns error in valid column %d", i)
		}
	}
}

func TestSetValue(t *testing.T) {
	var sudoku1 Sudoku
	var sudoku2 Sudoku

	// Must return error if value is below 1.
	for _, val := range [4]int{0, -10, -40, -100} {
		if err := sudoku1.SetValue(8, 8, val); err == nil {
			t.Errorf("Sudoku: Accepts value below 1: %d", val)
		}
	}

	// Must return error if value is above 9.
	for _, val := range [4]int{10, 100, 20, 34} {
		if err := sudoku1.SetValue(8, 8, val); err == nil {
			t.Errorf("Sudoku: Accepts value above 9: %d", val)
		}
	}

	// Must not return error for a value between 1 and 9.
	for i := 1; i <= 9; i++ {
		if err := sudoku1.SetValue(8, 8, i); err != nil {
			t.Errorf("Sudoku: Doesn't accept values between 1-9: %d", i)
		}
	}

	// Must return an error if we try to write in an invalid row.
	for _, val := range [4]int{-1, -2, 10, 11} {
		if err := sudoku1.SetValue(val, 0, 1); err == nil {
			t.Errorf("Sudoku: Invalid row %d", val)
		}
	}

	// Must return an error if we try to write in an invalid column.
	for _, val := range [4]int{-1, -2, -10, -100} {
		if err := sudoku1.SetValue(0, val, 1); err == nil {
			t.Errorf("Sudoku: Invalid column %d", val)
		}
	}

	// Must not return an error if we try to write in a valid row.
	for i := 0; i < 9; i++ {
		if err := sudoku1.SetValue(i, 0, 1); err != nil {
			t.Errorf("Sudoku: Returns error in valid row %d", i)
		}
	}

	// Must not return an error if we try to write in a valid column.
	for i := 0; i < 9; i++ {
		if err := sudoku1.SetValue(0, i, 1); err != nil {
			t.Errorf("Sudoku: Returns error in valid column %d", i)
		}
	}

	// Set some initial values on the sudoku.
	for i := 0; i < 9; i++ {
		sudoku2.SetInitialValue(i, i, 1)
	}

	// Must return an error when attempting to overwrite an initial value.
	for i := 0; i < 9; i++ {
		if err := sudoku2.SetValue(i, i, i); err == nil {
			t.Errorf("Sudoku: %v", err)
			t.Errorf("Sudoku: Can overwrite initial value on (%d, %d).", i, i)
		}
	}

	// Must not return an error when writting on non-initial value.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i != j {
				if err := sudoku2.SetValue(i, j, 1); err != nil {
					t.Errorf("Sudoku: %v", err)
					t.Errorf("Sudoku: Can't overwrite non-initial value on (%d, %d).", i, j)
				}
			}
		}
	}
}

func TestGetValue(t *testing.T) {
	var sudoku Sudoku
	var val int
	var err error

	// All initial values of the sudoku should be zero.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val, _ = sudoku.GetValue(i, j)

			if val != 0 {
				t.Errorf("Sudoku: Value should be %d but is %d", 0, val)
			}
		}
	}

	// Set initial values on the first row.
	for i := 0; i < 9; i++ {
		sudoku.SetInitialValue(0, i, i+1)
	}

	// Check that all the values on the first row can be obtained.
	for i := 0; i < 9; i++ {
		val, _ = sudoku.GetValue(0, i)

		if val != i+1 {
			t.Errorf("Sudoku: Value should be %d but is %d", i+1, val)
		}
	}

	// Set values on the last row.
	for i := 0; i < 9; i++ {
		sudoku.SetValue(8, i, i+1)
	}

	// Check that all values on the last row can be obtained.
	for i := 0; i < 9; i++ {
		val, _ = sudoku.GetValue(8, i)

		if val != i+1 {
			t.Errorf("Sudoku: Value should be %d but is %d", i+1, val)
		}
	}

	// Check that you can't obtain values on invalid rows.
	for _, val := range [5]int{-2, -10, 10, 100, 50} {
		_, err = sudoku.GetValue(val, 0)

		if err == nil {
			t.Errorf("Sudoku: Can obtain value of invalid row %d", val)
		}
	}

	// Check that you can't obtain values on invalid columns.
	for _, val := range [5]int{-2, -10, 10, 100, 50} {
		_, err = sudoku.GetValue(0, val)

		if err == nil {
			t.Errorf("Sudoku: Can obtain value of invalid column %d", val)
		}
	}
}

func TestGetRow(t *testing.T) {
	var sudoku Sudoku

	zero := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	one := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	valid := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// The initial state of a sudoku only contains zero's in all rows.
	for i := 0; i < 9; i++ {
		if sudoku.GetRow(i) != zero {
			t.Errorf("Sudoku: Row %d should be %v, but is %v", i, zero, sudoku.GetRow(i))
		}
	}

	// We fill the first row with ones.
	for i := 0; i < 9; i++ {
		sudoku.SetInitialValue(0, i, 1)
	}

	// The first row of the sudoku should contain only ones.
	if sudoku.GetRow(0) != one {
		t.Errorf("Sudoku: Row %d should be %v, but is %v", 0, one, sudoku.GetRow(0))
	}

	// We fill the last row with the sequence (1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < 9; i++ {
		sudoku.SetInitialValue(8, i, i+1)
	}

	// The last row of the sudoku should contain the values (1, 2, 3, 4, 5, 6, 7,
	// 8, 9).
	if sudoku.GetRow(8) != valid {
		t.Errorf("Sudoku: Row %d should be %v, but is %v", 8, valid, sudoku.GetRow(8))
	}
}

func TestGetColumn(t *testing.T) {
	var sudoku Sudoku

	zero := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	one := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	valid := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// The initial state of a sudoku only contains zero's in all columns.
	for i := 0; i < 9; i++ {
		if sudoku.GetColumn(i) != zero {
			t.Errorf("Sudoku: Column %d should be %v, but is %v", i, zero, sudoku.GetColumn(i))
		}
	}

	// We fill the first column with ones.
	for i := 0; i < 9; i++ {
		sudoku.SetInitialValue(i, 0, 1)
	}

	// The first column of the sudoku should contain only ones.
	if sudoku.GetColumn(0) != one {
		t.Errorf("Sudoku: Column %d should be %v, but is %v", 0, one, sudoku.GetColumn(0))
	}

	// We fill the last row with the sequence (1, 2, 3, 4, 5, 6, 7, 8, 9)
	for i := 0; i < 9; i++ {
		sudoku.SetInitialValue(i, 8, i+1)
	}

	// The last row of the sudoku should contain the values (1, 2, 3, 4, 5, 6, 7,
	// 8, 9).
	if sudoku.GetColumn(8) != valid {
		t.Errorf("Sudoku: Column %d should be %v, but is %v", 8, valid, sudoku.GetColumn(8))
	}
}

func TestGetBlock(t *testing.T) {
	var sudoku Sudoku

	zero := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	one := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	valid := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// The initial state of a sudoku only contains zero's in all blocks.
	for i := 0; i < 9; i++ {
		if sudoku.GetBlock(i) != zero {
			t.Errorf("Sudoku: Block %d should be %v, but is %v", i, zero, sudoku.GetBlock(i))
		}
	}

	// We fill the first block with ones.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sudoku.SetInitialValue(i, j, 1)
		}
	}

	// The first block of the sudoku should contain only ones.
	if sudoku.GetBlock(0) != one {
		t.Errorf("Sudoku: Block %d should be %v, but is %v", 0, one, sudoku.GetBlock(0))
	}

	// We fill the block 4 with the sequence (1, 2, 3, 4, 5, 6, 7, 8, 9)
	sudoku.SetInitialValue(3, 3, 1)
	sudoku.SetInitialValue(3, 4, 2)
	sudoku.SetInitialValue(3, 5, 3)
	sudoku.SetInitialValue(4, 3, 4)
	sudoku.SetInitialValue(4, 4, 5)
	sudoku.SetInitialValue(4, 5, 6)
	sudoku.SetInitialValue(5, 3, 7)
	sudoku.SetInitialValue(5, 4, 8)
	sudoku.SetInitialValue(5, 5, 9)

	// The block 4 of the sudoku should contain the values (1, 2, 3, 4, 5, 6, 7,
	// 8, 9).
	if sudoku.GetBlock(4) != valid {
		t.Errorf("Sudoku: Block %d should be %v, but is %v", 4, valid, sudoku.GetBlock(4))
	}

}

func TestIsValidRow(t *testing.T) {
	var sudoku Sudoku

	// Must return an error if any of the rows in the sudoku are valid.
	for i := 0; i < 9; i++ {
		if sudoku.IsValidRow(i) {
			t.Errorf("Sudoku: The following row is valid: %v", sudoku.GetRow(i))
		}
	}

	// Write only 1 in the first row.
	for i := 0; i < 9; i++ {
		sudoku.SetValue(0, i, 1)
	}

	// Must return an error if the first row is valid.
	if sudoku.IsValidRow(0) {
		t.Errorf("Sudoku: The following row is valid: %v", sudoku.GetRow(0))
	}

	// Write valid rows in all the sudoku.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku.SetValue(i, j, j+1)
		}
	}

	// Must not return an error as all the rows are now valid.
	for i := 0; i < 9; i++ {
		if !sudoku.IsValidRow(i) {
			t.Errorf("Sudoku: The following row is invalid: %v", sudoku.GetRow(i))
		}
	}
}

func TestIsValidColumn(t *testing.T) {
	var sudoku Sudoku

	// Must return an error if any of the columns in the sudoku are valid.
	for i := 0; i < 9; i++ {
		if sudoku.IsValidColumn(i) {
			t.Errorf("Sudoku: The following column is valid: %v", sudoku.GetColumn(i))
		}
	}

	// Write only 1 in the first column.
	for i := 0; i < 9; i++ {
		sudoku.SetValue(i, 0, 1)
	}

	// Must return an error if the first column is valid.
	if sudoku.IsValidColumn(0) {
		t.Errorf("Sudoku: The following column is valid: %v", sudoku.GetColumn(0))
	}

	// Write valid columns in all the sudoku.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku.SetValue(j, i, j+1)
		}
	}

	// Must not return an error as all the columns are now valid.
	for i := 0; i < 9; i++ {
		if !sudoku.IsValidColumn(i) {
			t.Errorf("Sudoku: The following row is invalid: %v", sudoku.GetColumn(i))
		}
	}
}

func TestIsValidBlock(t *testing.T) {
	var sudoku Sudoku

	// Must return an error if any of the columns in the sudoku are valid.
	for i := 0; i < 9; i++ {
		if sudoku.IsValidBlock(i) {
			t.Errorf("Sudoku: The following block is valid: %v", sudoku.GetBlock(i))
		}
	}

	// Write only 1 in the first block.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sudoku.SetValue(i, j, 1)
		}
	}

	// Must return an error if the first column is valid.
	if sudoku.IsValidBlock(0) {
		t.Errorf("Sudoku: The following column is valid: %v", sudoku.GetColumn(0))
	}

	// Write a valid block in block 1.
	sudoku.SetValue(0, 0, 1)
	sudoku.SetValue(0, 1, 2)
	sudoku.SetValue(0, 2, 3)
	sudoku.SetValue(1, 0, 4)
	sudoku.SetValue(1, 1, 5)
	sudoku.SetValue(1, 2, 6)
	sudoku.SetValue(2, 0, 7)
	sudoku.SetValue(2, 1, 8)
	sudoku.SetValue(2, 2, 9)

	// Must not return an error as the block 1 is now valid.
	if !sudoku.IsValidBlock(0) {
		t.Errorf("Sudoku: The following block is invalid: %v", sudoku.GetBlock(0))
	}

	// Write another valid block in block 4.
	sudoku.SetValue(3, 0, 9)
	sudoku.SetValue(3, 1, 8)
	sudoku.SetValue(3, 2, 7)
	sudoku.SetValue(4, 0, 6)
	sudoku.SetValue(4, 1, 3)
	sudoku.SetValue(4, 2, 5)
	sudoku.SetValue(5, 0, 2)
	sudoku.SetValue(5, 1, 4)
	sudoku.SetValue(5, 2, 1)

	// Must not return an error as the block 4 is now valid.
	if !sudoku.IsValidBlock(0) {
		t.Errorf("Sudoku: The following block is invalid: %v", sudoku.GetBlock(4))
	}

	// We modifiy block 1 to not be valid anymore.
	sudoku.SetValue(2, 2, 9)

	// Must return an error as the block 1 is no longer valid.
	if !sudoku.IsValidBlock(0) {
		t.Errorf("Sudoku: The following block is invalid: %v", sudoku.GetBlock(0))
	}
}

func TestIsComplete(t *testing.T) {
	var sudoku Sudoku

	// Must return error if the empty sudoky is considered completed.
	if sudoku.IsComplete() {
		t.Errorf("Sudoku: The following sudoku is considered to be complete when its not %v", sudoku.ToString())
	}

	// We fill the sudoku with ones.
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sudoku.SetValue(i, j, 1)
		}
	}

	// Must return error if a sudoku full of ones is considered completed.
	if sudoku.IsComplete() {
		t.Errorf("Sudoku: The following sudoku is considered to be complete when its not %v", sudoku.ToString())
	}

	// The following Sudoku is considered complete.
	// ╔───┬───┬───╦───┬───┬───╦───┬───┬───╗
	// │ 6 │ 3 │ 8 │ 2 │ 4 │ 7 │ 1 │ 9 │ 5 │
	// ├───┼───┼───┼───┼───┼───┼───┼───┼───┤
	// │ 4 │ 2 │ 5 │ 8 │ 1 │ 9 │ 7 │ 6 │ 3 │
	// ├───┼───┼───┼───┼───┼───┼───┼───┼───┤
	// │ 7 │ 9 │ 1 │ 5 │ 3 │ 6 │ 4 │ 8 │ 2 │
	// ╠───┼───┼───╬───┼───┼───╬───┼───┼───╣
	// │ 9 │ 5 │ 7 │ 1 │ 8 │ 3 │ 6 │ 2 │ 4 │
	// ├───┼───┼───┼───┼───┼───┼───┼───┼───┤
	// │ 3 │ 8 │ 2 │ 7 │ 6 │ 4 │ 9 │ 5 │ 1 │
	// ├───┼───┼───┼───┼───┼───┼───┼───┼───┤
	// │ 1 │ 6 │ 4 │ 9 │ 5 │ 2 │ 3 │ 7 │ 8 │
	// ╠───┼───┼───╬───┼───┼───╬───┼───┼───╣
	// │ 2 │ 4 │ 9 │ 3 │ 7 │ 5 │ 8 │ 1 │ 6 │
	// ├───┼───┼───┼───┼───┼───┼───┼───┼───┤
	// │ 8 │ 7 │ 6 │ 4 │ 2 │ 1 │ 5 │ 3 │ 9 │
	// ├───┼───┼───┼───┼───┼───┼───┼───┼───┤
	// │ 5 │ 1 │ 3 │ 6 │ 9 │ 8 │ 2 │ 4 │ 7 │
	// ╚───┴───┴───╩───┴───┴───╩───┴───┴───╝

	// We fill block zero.
	sudoku.SetValue(0, 0, 6)
	sudoku.SetValue(0, 1, 3)
	sudoku.SetValue(0, 2, 8)
	sudoku.SetValue(1, 0, 4)
	sudoku.SetValue(1, 1, 2)
	sudoku.SetValue(1, 2, 5)
	sudoku.SetValue(2, 0, 7)
	sudoku.SetValue(2, 1, 9)
	sudoku.SetValue(2, 2, 1)

	// We fill block one.
	sudoku.SetValue(0, 3, 2)
	sudoku.SetValue(0, 4, 4)
	sudoku.SetValue(0, 5, 7)
	sudoku.SetValue(1, 3, 8)
	sudoku.SetValue(1, 4, 1)
	sudoku.SetValue(1, 5, 9)
	sudoku.SetValue(2, 3, 5)
	sudoku.SetValue(2, 4, 3)
	sudoku.SetValue(2, 5, 6)

	// We fill block two.
	sudoku.SetValue(0, 6, 1)
	sudoku.SetValue(0, 7, 9)
	sudoku.SetValue(0, 8, 5)
	sudoku.SetValue(1, 6, 7)
	sudoku.SetValue(1, 7, 6)
	sudoku.SetValue(1, 8, 3)
	sudoku.SetValue(2, 6, 4)
	sudoku.SetValue(2, 7, 8)
	sudoku.SetValue(2, 8, 2)

	// We fill block three.
	sudoku.SetValue(3, 0, 9)
	sudoku.SetValue(3, 1, 5)
	sudoku.SetValue(3, 2, 7)
	sudoku.SetValue(4, 0, 3)
	sudoku.SetValue(4, 1, 8)
	sudoku.SetValue(4, 2, 2)
	sudoku.SetValue(5, 0, 1)
	sudoku.SetValue(5, 1, 6)
	sudoku.SetValue(5, 2, 4)

	// We fill block four.
	sudoku.SetValue(3, 3, 1)
	sudoku.SetValue(3, 4, 8)
	sudoku.SetValue(3, 5, 3)
	sudoku.SetValue(4, 3, 7)
	sudoku.SetValue(4, 4, 6)
	sudoku.SetValue(4, 5, 4)
	sudoku.SetValue(5, 3, 9)
	sudoku.SetValue(5, 4, 5)
	sudoku.SetValue(5, 5, 2)

	// We fill block five.
	sudoku.SetValue(3, 6, 6)
	sudoku.SetValue(3, 7, 2)
	sudoku.SetValue(3, 8, 4)
	sudoku.SetValue(4, 6, 9)
	sudoku.SetValue(4, 7, 5)
	sudoku.SetValue(4, 8, 1)
	sudoku.SetValue(5, 6, 3)
	sudoku.SetValue(5, 7, 7)
	sudoku.SetValue(5, 8, 8)

	// We fill block six.
	sudoku.SetValue(6, 0, 2)
	sudoku.SetValue(6, 1, 4)
	sudoku.SetValue(6, 2, 9)
	sudoku.SetValue(7, 0, 8)
	sudoku.SetValue(7, 1, 7)
	sudoku.SetValue(7, 2, 6)
	sudoku.SetValue(8, 0, 5)
	sudoku.SetValue(8, 1, 1)
	sudoku.SetValue(8, 2, 3)

	// We fill block seven.
	sudoku.SetValue(6, 3, 3)
	sudoku.SetValue(6, 4, 7)
	sudoku.SetValue(6, 5, 5)
	sudoku.SetValue(7, 3, 4)
	sudoku.SetValue(7, 4, 2)
	sudoku.SetValue(7, 5, 1)
	sudoku.SetValue(8, 3, 6)
	sudoku.SetValue(8, 4, 9)
	sudoku.SetValue(8, 5, 8)

	// We fill block eigth.
	sudoku.SetValue(6, 6, 8)
	sudoku.SetValue(6, 7, 1)
	sudoku.SetValue(6, 8, 6)
	sudoku.SetValue(7, 6, 5)
	sudoku.SetValue(7, 7, 3)
	sudoku.SetValue(7, 8, 9)
	sudoku.SetValue(8, 6, 2)
	sudoku.SetValue(8, 7, 4)
	sudoku.SetValue(8, 8, 7)

	// Every row has to be valid.
	for i := 0; i < 9; i++ {
		if !sudoku.IsValidRow(i) {
			t.Errorf("Sudoku: Following row should be valid %v", sudoku.GetRow(i))
		}
	}

	// Every column has to be valid.
	for i := 0; i < 9; i++ {
		if !sudoku.IsValidColumn(i) {
			t.Errorf("Sudoku: Following column should be valid %v", sudoku.GetColumn(i))
		}
	}

	// Every block has to be valid.
	for i := 0; i < 9; i++ {
		if !sudoku.IsValidBlock(i) {
			t.Errorf("Sudoku: Following block should be valid %v", sudoku.GetBlock(i))
		}
	}

	// Must not return false as the given sudoku is complete.

	if !sudoku.IsComplete() {
		t.Errorf("Sudoku: The following Sudoku returns false when asked if complete: \n%v", sudoku.ToString())
	}
}
