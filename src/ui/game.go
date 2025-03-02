package ui

import (
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/maucarrui/Sudoku2Go/internal"
	"github.com/maucarrui/Sudoku2Go/src/model"
	"strconv"
)

type Sudoku = model.Sudoku
type SudokuError = model.SudokuError

type StringSet = internal.StringSet

type GameState int64

type Game struct {
	cursor_x      int
	cursor_y      int
	initialSudoku *Sudoku
	sudoku        *Sudoku
	error         *SudokuError
	message       string
	timer         stopwatch.Model

	// User Interface specifics.
	blinking bool
}

func NewGame() Game {
	nonCompleteSudoku := []int{
		6, 0, 0, 0, 7, 0, 0, 1, 5,
		0, 0, 1, 0, 0, 0, 8, 6, 4,
		5, 0, 4, 0, 0, 0, 7, 0, 3,
		9, 0, 3, 0, 0, 1, 4, 0, 0,
		0, 0, 0, 4, 3, 5, 6, 0, 0,
		0, 4, 6, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 6, 2, 0, 5, 7,
		3, 0, 0, 0, 0, 0, 2, 8, 0,
		2, 0, 0, 0, 5, 8, 0, 0, 6}

	i, _ := model.NewSudoku(nonCompleteSudoku)
	s, _ := model.NewSudoku(nonCompleteSudoku)

	return Game{
		cursor_x:      0,
		cursor_y:      0,
		initialSudoku: i,
		sudoku:        s,
		error:         nil,
		message:       "",
		timer:         stopwatch.New(),
	}
}

func (game *Game) MoveCursor(direction string) {
	// Keys that move the cursor up, down, left, and right.
	upKeys := internal.NewStringSetFromValues(
		[]string{"ctrl+p", "up", "w"},
	)

	downKeys := internal.NewStringSetFromValues(
		[]string{"ctrl+n", "down", "s"},
	)

	leftKeys := internal.NewStringSetFromValues(
		[]string{"ctrl+b", "left", "a"},
	)

	rightKeys := internal.NewStringSetFromValues(
		[]string{"ctrl+f", "right", "d"},
	)

	if upKeys.Contains(direction) {
		// Move up.
		if game.cursor_y > 0 {
			game.cursor_y--
		}
	} else if downKeys.Contains(direction) {
		// Move down.
		if game.cursor_y < 8 {
			game.cursor_y++
		}
	} else if leftKeys.Contains(direction) {
		// Move left.
		if game.cursor_x > 0 {
			game.cursor_x--
		}
	} else if rightKeys.Contains(direction) {
		// Move right.
		if game.cursor_x < 8 {
			game.cursor_x++
		}
	}

	game.error = nil
}

func (game *Game) AddValue(key string) {
	// Parse the string into an integer.
	value, _ := strconv.Atoi(key)

	row := game.cursor_y
	col := game.cursor_x

	if val, _ := game.initialSudoku.GetValue(row, col); val != 0 {
		game.message = "Can't overwrite initial value"
		return
	} else if val, _ := game.sudoku.GetValue(row, col); val != 0 {
		game.sudoku.RemoveValue(row, col)
	}

	game.error = game.sudoku.SetValue(row, col, value)
}

func (game *Game) RemoveValue() {
	row := game.cursor_y
	col := game.cursor_x

	if val, _ := game.initialSudoku.GetValue(row, col); val != 0 {
		return
	}

	game.error = game.sudoku.RemoveValue(row, col)
}

func (game *Game) IsComplete() bool {
	return game.sudoku.IsComplete()
}
