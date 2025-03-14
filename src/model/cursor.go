package model

// The Cursor keeps track of the position in the Sudoku game.
type Cursor struct {
	Row int
	Col int
}

// Returns a new cursor located at the top-left of the board.
func NewCursor() *Cursor {
	return &Cursor{
		Row: 0,
		Col: 0,
	}
}

// If possible, moves the cursor to the left.
func (cursor *Cursor) MoveLeft() {
	if cursor.Col > 0 {
		cursor.Col--
	}
}

// If possible, moves the cursor to the right.
func (cursor *Cursor) MoveRight() {
	if cursor.Col < 8 {
		cursor.Col++
	}
}

// If possible, moves the cursor up.
func (cursor *Cursor) MoveUp() {
	if cursor.Row > 0 {
		cursor.Row--
	}
}

// If possible, moves the cursor down.
func (cursor *Cursor) MoveDown() {
	if cursor.Row < 8 {
		cursor.Row++
	}
}
