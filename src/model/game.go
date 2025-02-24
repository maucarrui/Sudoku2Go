package model

// The sudoku game.
type Game struct {

	// The initial sudoku, its values can't be modified by the user.
	initialSudoku *Sudoku

	// A copy of the initial sudoku, this sudoku is the one the user modifies.
	sudoku *Sudoku

	// The cursor's coordinates.
	cursorX int
	cursorY int
}
