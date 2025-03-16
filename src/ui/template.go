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
	gridBuilder.WriteString("\n")

	return gridBuilder.String()
}

// Returns the upper part of the sudoku grid.
func upperGridToString() string {
	return gridToString("╔══╦╗")
}

// Returns the middle part of the sudoku grid.
func middleGridToString() string {
	return gridToString("║─┼║║")
}

// Returns the middle (delimited) part of the sudoku grid.
func middleDelimitedGridToString() string {
	return gridToString("╠══╬╣")
}

// Returns the lower part of the sudoku grid.
func lowerGridToString() string {
	return gridToString("╚══╩╝")
}

func isCursorOnValue(game Game, row, col int) bool {
	return (row == game.cursor.Row) && (col == game.cursor.Col)
}

func isConflictingValue(game Game, row, col int) bool {
	return game.error != nil &&
		game.error.GetConflictRow() == row &&
		game.error.GetConflictColumn() == col
}

func isMatchingValue(game Game, row, col int) bool {
	selectedValue, _ := game.sudoku.GetValue(game.cursor.Row, game.cursor.Col)
	currentValue, _ := game.sudoku.GetValue(row, col)
	return selectedValue == currentValue
}

func isInitialValue(game Game, row, col int) bool {
	value, _ := game.initialSudoku.GetValue(row, col)
	return value != 0
}

func isPlacedValue(game Game, row, col int) bool {
	initialValue, _ := game.initialSudoku.GetValue(row, col)
	placedValue, _ := game.sudoku.GetValue(row, col)

	return initialValue == 0 && placedValue != 0
}

func isEmptyValue(game Game, row, col int) bool {
	value, _ := game.sudoku.GetValue(row, col)
	return value == 0
}

func selectedValueToString(game Game, row, col int) string {
	value, _ := game.sudoku.GetValue(row, col)
	str := strconv.Itoa(value)
	if game.blinking {
		return selectedValueStyleBlinking(str)
	} else {
		return selectedValueStyle(str)
	}
}

func cursorToString(game Game) string {
	if game.blinking {
		return selectedValueStyleBlinking("█")
	} else {
		return selectedValueStyle("█")
	}
}

func selectionToString(game Game, row, col int) string {
	value, _ := game.sudoku.GetValue(row, col)
	if value != 0 {
		return selectedValueToString(game, row, col)
	} else {
		return cursorToString(game)
	}
}

func conflictValueToString(game Game, row, col int) string {
	value, _ := game.sudoku.GetValue(row, col)
	str := strconv.Itoa(value)
	return conflictValueStyle(str)
}

func matchingValueToString(game Game, row, col int) string {
	value, _ := game.sudoku.GetValue(row, col)
	if value != 0 {
		str := strconv.Itoa(value)
		return matchingValuesStyle(str)
	} else {
		return " "
	}
}

func initialValueToString(game Game, row, col int) string {
	value, _ := game.sudoku.GetValue(row, col)
	str := strconv.Itoa(value)
	return initialValueStyle(str)
}

func placedValueToString(game Game, row, col int) string {
	value, _ := game.sudoku.GetValue(row, col)
	str := strconv.Itoa(value)
	return placedValuesStyle(str)
}

func emptyValueToString() string {
	return " "
}

func paintValue(game Game, row, col int) string {
	if isCursorOnValue(game, row, col) {
		return selectionToString(game, row, col)
	} else if isConflictingValue(game, row, col) {
		return conflictValueToString(game, row, col)
	} else if isMatchingValue(game, row, col) {
		return matchingValueToString(game, row, col)
	} else if isInitialValue(game, row, col) {
		return initialValueToString(game, row, col)
	} else if isPlacedValue(game, row, col) {
		return placedValueToString(game, row, col)
	} else if isEmptyValue(game, row, col) {
		return emptyValueToString()
	}

	return ""
}

func rowToString(game Game, row int) string {
	rowStr := "║"

	for col := 0; col < 9; col++ {

		rowStr += " "
		rowStr += paintValue(game, row, col)
		rowStr += " "

		if (col+1)%3 == 0 {
			rowStr += "║"
		} else {
			rowStr += "│"
		}
	}

	return rowStr + "\n"
}

func gameToString(game Game) string {
	str := ""

	for row := 0; row < 9; row++ {

		if row == 0 {
			str += upperGridToString()
		} else if row%3 == 0 {
			str += middleDelimitedGridToString()
		} else {
			str += middleGridToString()
		}

		str += rowToString(game, row)
	}

	str += lowerGridToString()
	return str
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
