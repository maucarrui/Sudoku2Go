package ui

import (
	"math"
	"strconv"
	"strings"
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

func gridToString(gridStructure string) string {
	// Cast the string to a rune array for easier UTF-8 handling.
	runes := []rune(gridStructure)

	start := runes[0]
	number := runes[1]
	numberSeparator := runes[2]
	blockSeparator := runes[3]
	end := runes[4]

	var gridBuilder strings.Builder
	gridBuilder.WriteRune(start)
	for i := 0; i < 9; i++ {

		// Each number takes a total of 3 characters.
		for j := 0; j < 3; j++ {
			gridBuilder.WriteRune(number)
		}

		if (i + 1) == 9 {
			gridBuilder.WriteRune(end)
		} else if (i+1)%3 == 0 {
			gridBuilder.WriteRune(blockSeparator)
		} else {
			gridBuilder.WriteRune(numberSeparator)
		}
	}

	return gridBuilder.String()
}

// Returns the upper part of the sudoku grid.
func upperGridToString() string {
	return gridToString("╔══╦╗")
}

// Returns the middle part of the sudoku grid.
func middleGridToString() string {
	return gridToString("║─┼║╣")
}

// Returns the middle (delimited) part of the sudoku grid.
func middleDelimitedGridToString() string {
	return gridToString("╠══╬╣")
}

// Returns the lower part of the sudoku grid.
func lowerGridToString() string {
	return gridToString("╚══╩╝")
}

// Returns the given sudoku in String format.
func PrintGame(game Game) string {
	cursor_x := game.cursor_x
	cursor_y := game.cursor_y
	sudoku := game.sudoku
	initialSudoku := game.initialSudoku
	error := game.error
	blinking := game.blinking

	conflict_x := -1
	conflict_y := -1
	if error != nil {
		conflict_x = error.GetConflictColumn()
		conflict_y = error.GetConflictRow()
	}

	selectedValue, _ := sudoku.GetValue(cursor_y, cursor_x)

	// Draw the upper part of the grid first.
	sudokuString := upperGridToString() + "\n"

	for i, row := range sudoku.GetValues() {

		// Print the current row values.
		for j := 0; j < len(row); j++ {

			if (j == 0) || (j%3 == 0) {
				sudokuString += "║ "
			} else {
				sudokuString += "│ "
			}

			value := row[j]
			strValue := strconv.Itoa(row[j])

			if (cursor_x == j) && (cursor_y == i) {
				if row[j] != 0 {
					if blinking {
						sudokuString += selectedValueStyleBlinking(strValue)
					} else {
						sudokuString += selectedValueStyle(strValue)
					}
				} else {
					if blinking {
						sudokuString += selectedValueStyleBlinking("█")
					} else {
						sudokuString += selectedValueStyle("█")
					}
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

		if i+1 == 9 {
			sudokuString += lowerGridToString() + "\n"
		} else if (i+1)%3 == 0 {
			sudokuString += middleDelimitedGridToString() + "\n"
		} else {
			sudokuString += middleGridToString() + "\n"
		}
	}

	return sudokuString
}

func BlinkingEnabledToString(game Game) string {
	str := "(Blinking: "
	if game.blinking {
		str += blinkingEnabledStyle("Enabled") + " )"
	} else {
		str += blinkingDisabledStyle("Disabled") + ")"
	}
	return str
}

func RulesToString() string {
	rules := "Rules: \n\n"
	rules += "A Sudoku is 9x9 board whose values range from 1 to 9. \n\n"
	rules += "Each row, column, and 3x3 square contained inside "
	rules += "the board shall have the number 1 through 9. \n\n"
	rules += "No repetitions are allowed."

	return rules
}

func ControlsToString(game Game) string {
	controls := "Controls: \n\n"
	controls += "Place a number using the 1-9 keys.\n\n"
	controls += "Up:    Ctrl+p, Up Arrow \n"
	controls += "Down:  Ctrl+n, Down Arrow \n"
	controls += "Left:  Ctrl+b, Left Arrow \n"
	controls += "Right: Ctrl+f, Right Arrow \n\n"
	controls += "Toggle blinking cursor: t\n"
	controls += "  " + BlinkingEnabledToString(game) + "\n\n"
	controls += "Quit:  Ctrl+q, q, Escape \n"

	return controls
}
