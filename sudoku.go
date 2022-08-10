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
	if val >= 1 && val <= 9 {
		sudoku.values[x][y] = val
		sudoku.initialValues[x][y] = val
	} else {
		return errors.New("Sudoku: Not a valid entry.")
	}

	return nil
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
