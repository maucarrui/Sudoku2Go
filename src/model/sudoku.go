package main

import (
	"errors"  // Error handling.
	"strconv" // String Conversions. (Integer to String)
)

type Sudoku struct {
	// All the sudoku values.
	values [9][9]int

	// The initial sudoku values; you can't modify this ones while playing.
	initialValues [9][9]int
}

// Set an initial value for the sudoku in the cell on the row x and column
// y. These values can't be modified while playing. Must be between 1 and 9,
// otherwise an error is returned.
func (sudoku *Sudoku) SetInitialValue(x, y, val int) error {
	if x < 0 || x > 8 {
		return errors.New("Sudoku: Invalid row.")
	}

	if y < 0 || y > 8 {
		return errors.New("Sudoku: Invalid column.")
	}

	if val >= 1 && val <= 9 {
		sudoku.values[x][y] = val
		sudoku.initialValues[x][y] = val
	} else {
		return errors.New("Sudoku: Not a valid entry.")
	}

	return nil
}

// Set a value for the sudoku in the cell on the row x and column y. Must be
// between 1 and 9, otherwise an error is returned.
func (sudoku *Sudoku) SetValue(x, y, val int) error {
	if x < 0 || x > 8 {
		return errors.New("Sudoku: Invalid row.")
	}

	if y < 0 || y > 8 {
		return errors.New("Sudoku: Invalid column.")
	}

	if sudoku.initialValues[x][y] != 0 {
		return errors.New("Sudoku: Can't overwrite initial value.")
	}

	if val >= 1 && val <= 9 {
		sudoku.values[x][y] = val
	} else {
		return errors.New("Sudoku: Not a valid entry.")
	}

	return nil
}

// Returns the value of the sudoku on the row x and column y.
func (sudoku *Sudoku) GetValue(x, y int) (int, error) {
	if x < 0 || x > 8 {
		return 0, errors.New("Sudoku: Invalid row.")
	}

	if y < 0 || y > 8 {
		return 0, errors.New("Sudoku: Invalid column.")
	}

	return sudoku.values[x][y], nil
}

// Returns the row x of the sudoku as an array of size 9.
func (sudoku *Sudoku) GetRow(x int) [9]int {
	var row [9]int

	for i := 0; i < 9; i++ {
		row[i] = sudoku.values[x][i]
	}

	return row
}

// Returns the column y of the sudoku as an array of size 9.
func (sudoku *Sudoku) GetColumn(x int) [9]int {
	var column [9]int

	for i := 0; i < 9; i++ {
		column[i] = sudoku.values[i][x]
	}

	return column
}

// A sudoku consists of 9 blocks, which are 3x3 matrices. The following diagram
// indicates the enumeration of each block.
// ╔───┬───┬───╦───┬───┬───╦───┬───┬───╗
// │           │           │           │
// ├           ┼           ┼          ─┤
// │     0     │     1     │     2     │
// ├           ┼           ┼          ─┤
// │           │           │           │
// ╠───┼───┼───╬───┼───┼───╬───┼───┼───╣
// │           │           │           │
// ├           ┼           ┼           ┤
// │     3     │     4     │     5     │
// ├           ┼           ┼           ┤
// │           │           │           │
// ╠───┼───┼───╬───┼───┼───╬───┼───┼───╣
// │           │           │           │
// ├           ┼           ┼           ┤
// │     6     │     7     │     8     │
// ├           ┼           ┼           ┤
// │           │           │           │
// ╚───┴───┴───╩───┴───┴───╩───┴───┴───╝
// GetBlock returns the content of the the given block as an array of size 9.
func (sudoku *Sudoku) GetBlock(z int) [9]int {
	var block [9]int
	var row, col int

	if z > 5 {
		row = 2
	} else if z > 2 {
		row = 1
	} else {
		row = 0
	}

	col = z % 3

	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			block[k] = sudoku.values[(row*3)+i][(col*3)+j]
			k++
		}
	}

	return block
}

// Returns true if the row x does not contain any repeated values and each value
// is between 1 and 9. The sum of a valid row will always be (10)(9)/2 = 45, and
// this value can only be obtained if each value is diferent and between 1 and
// 9.
func (sudoku *Sudoku) IsValidRow(x int) bool {
	sum := 0
	val := 0

	for i := 0; i < 9; i++ {
		val = sudoku.values[x][i]

		if val < 1 || val > 9 {
			return false
		}

		sum += val
	}

	return sum == 45
}

// Returns true if the column y does not contain any repeated values and each
// value is between 1 and 9. The sum of a valid column will always be (10)(9)/2
// = 45, and this value can only be obtained if each value is diferent and
// between 1 and 9.
func (sudoku *Sudoku) IsValidColumn(y int) bool {
	sum := 0
	val := 0

	for i := 0; i < 9; i++ {
		val = sudoku.values[i][y]

		if val < 1 || val > 9 {
			return false
		}

		sum += val
	}

	return sum == 45
}

// Returns true if the block z does not contain any repeated values and each
// value is between 1 and 9. The sum of a valid block will always be (10)(9)/2 =
// 45, and this value can only be obtained if each value is diferent and between
// 1 and 9. Reference of the enumerations of blocks on method @GetBlock.
func (sudoku *Sudoku) IsValidBlock(z int) bool {
	var row, col int

	if z > 5 {
		row = 2
	} else if z > 2 {
		row = 1
	} else {
		row = 0
	}

	col = z % 3

	sum := 0
	val := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val = sudoku.values[(row*3)+i][(col*3)+j]

			if val < 1 || val > 9 {
				return false
			}

			sum += val
		}
	}

	return sum == 45
}

// A completed sudoku is a Sudoku where all its rows, columns, and blocks are
// valid.
func (sudoku *Sudoku) IsComplete() bool {
	for i := 0; i < 9; i++ {
		if !sudoku.IsValidRow(i) ||
			!sudoku.IsValidBlock(i) ||
			!sudoku.IsValidBlock(i) {
			return false
		}
	}

	return true
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
