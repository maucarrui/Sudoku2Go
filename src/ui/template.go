package ui

import (
	"math"
	"strconv"
	"time"
)

type Duration = time.Duration

// Returns a string representation of the given time unit.
// If the time unit is less than "10", a zero is appended to the beginning
// of the string.
func timeUnitToString(timeUnit int) string {
	str := ""
	if timeUnit < 10 {
		str += "0"
	}
	return str + strconv.Itoa(timeUnit)
}

func PrintElapsedTime(elapsedTime Duration) string {
	elapsedSeconds := elapsedTime.Seconds()

	seconds := int(math.Mod(elapsedSeconds, 60))
	minutes := int(math.Mod(elapsedSeconds/60, 60))
	hours := int((elapsedSeconds / 60) / 60)

	ss := timeUnitToString(seconds)
	mm := timeUnitToString(minutes)
	hh := timeUnitToString(hours)

	return hh + ":" + mm + ":" + ss
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
			gridString = "║"
		}
	case 3:
		gridString = "╚"
	}

	for i := 0; i < length; i++ {
		if delim {
			gridString += "═══"
		} else {
			gridString += "───"
		}

		if i < (length - 1) {
			switch part {
			case 1:
				if (i+1)%3 == 0 {
					gridString += "╦"
				} else {
					gridString += "═"
				}

			case 2:
				if delim && (i+1)%3 == 0 {
					gridString += "╬"
				} else if (i+1)%3 == 0 {
					gridString += "║"
				} else if delim {
					gridString += "═"
				} else {
					gridString += "┼"
				}

			case 3:
				if (i+1)%3 == 0 {
					gridString += "╩"
				} else {
					gridString += "═"
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
					gridString += "║"
				}
			case 3:
				gridString += "╝"
			}
		}
	}

	return gridString
}

// Returns the given sudoku in String format.
func PrintGame(game Game) string {
	var delim bool

	cursor_x := game.cursor_x
	cursor_y := game.cursor_y
	sudoku := game.sudoku
	initialSudoku := game.initialSudoku
	error := game.error

	conflict_x := -1
	conflict_y := -1
	if error != nil {
		conflict_x = error.GetConflictColumn()
		conflict_y = error.GetConflictRow()
	}

	selectedValue, _ := sudoku.GetValue(cursor_y, cursor_x)
	sudokuString := printGridPart(1, 9, true) + "\n"

	for i, row := range sudoku.GetValues() {

		// Print the current row values.
		for j := 0; j < len(row); j++ {

			if (i+1)%3 == 0 {
				delim = true
			} else {
				delim = false
			}

			if (j == 0) || (j%3 == 0) {
				sudokuString += "║ "
			} else {
				sudokuString += "│ "
			}

			value := row[j]
			strValue := strconv.Itoa(row[j])

			if (cursor_x == j) && (cursor_y == i) {
				if row[j] != 0 {
					sudokuString += selectedValueStyle(strValue)
				} else {
					sudokuString += selectedValueStyle("█")
				}
			} else if (conflict_x == j) && (conflict_y == i) {
				sudokuString += conflictValueStyle(strValue)
			} else if selectedValue == value && value != 0 {
				sudokuString += matchingValuesStyle(strValue)
			} else if value != 0 {
				if initialValue, _ := initialSudoku.GetValue(i, j); initialValue != 0 {
					sudokuString += initialValueStyle(strValue)
				} else {
					sudokuString += placedValuesStyle(strValue)
				}

			} else {
				sudokuString += " "
			}

			sudokuString += " "

			if j == len(row)-1 {
				sudokuString += "║\n"
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

func PrintInstructions() string {

	instructions := "Instructions: \n\n"
	instructions += "A Sudoku is 9x9 board whose values range from 1 to 9. \n\n"
	instructions += "Each row, column, and 3x3 square contained inside "
	instructions += "the board should have the number 1 through 9. \n\n"
	instructions += "No repetitions are allowed."

	controls := "Controls: \n\n"
	controls += "Up:    Ctrl+p, Up Arrow \n"
	controls += "Down:  Ctrl+n, Down Arrow \n"
	controls += "Left:  Ctrl+b, Left Arrow \n"
	controls += "Right: Ctrl+f, Right Arrow \n"
	controls += "Quit:  Ctrl+q, q \n"

	return instructions + "\n\n" + controls
}
