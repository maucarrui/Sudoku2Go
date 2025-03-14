package model

import (
	"testing"
)

// Tests the constructor function for the Cursor struct.
func TestNewCursor(t *testing.T) {
	errMsg := "Cursor.NewCursor: "
	errMsg += "Expected row and col to be 0. "
	errMsg += "row: %d, col %d."

	cursor := NewCursor()
	if cursor.Row != 0 && cursor.Col != 0 {
		t.Errorf(errMsg, cursor.Row, cursor.Col)
	}
}

// Tests the move left function of the cursor.
func TestMoveLeft(t *testing.T) {
	errMsg := "Cursor.MoveLeft: "
	errMsg += "Expected col to be %d after moving to the left, "
	errMsg += "but got %d."

	cursor := NewCursor()
	cursor.Col = 8
	expectedCol := 8

	for i := 0; i < 8; i++ {
		expectedCol--
		cursor.MoveLeft()

		if cursor.Col != expectedCol {
			t.Errorf(errMsg, expectedCol, cursor.Col)
		}
	}
}

// Tests the move left function of the cursor,
// when the cursor is at the left border (column 0).
func TestMoveLeftAtBorder(t *testing.T) {
	errMsg := "Cursor.MoveLeft: "
	errMsg += "Expected col to not change when moving to the left, "
	errMsg += "but got %d."

	cursor := NewCursor()
	cursor.MoveLeft()

	if cursor.Col != 0 {
		t.Errorf(errMsg, cursor.Col)
	}
}

// Tests the move right function of the cursor.
func TestMoveRight(t *testing.T) {
	errMsg := "Cursor.MoveRight: "
	errMsg += "Expected col to be %d after moving to the right, "
	errMsg += "but got %d."

	cursor := NewCursor()
	expectedCol := 0

	for i := 0; i < 8; i++ {
		expectedCol++
		cursor.MoveRight()

		if cursor.Col != expectedCol {
			t.Errorf(errMsg, expectedCol, cursor.Col)
		}
	}
}

// Tests the move right function of the cursor,
// when the cursor is at the right border (column 8).
func TestMoveRightAtBorder(t *testing.T) {
	errMsg := "Cursor.MoveRight: "
	errMsg += "Expected col to not change when moving to the right, "
	errMsg += "but got %d."

	cursor := NewCursor()
	cursor.Col = 8
	cursor.MoveRight()

	if cursor.Col != 8 {
		t.Errorf(errMsg, cursor.Col)
	}
}

// Tests the move up function of the cursor.
func TestMoveUp(t *testing.T) {
	errMsg := "Cursor.MoveUp: "
	errMsg += "Expected row to be %d after moving up, "
	errMsg += "but got %d."

	cursor := NewCursor()
	cursor.Row = 8
	expectedRow := 8

	for i := 0; i < 8; i++ {
		expectedRow--
		cursor.MoveUp()

		if cursor.Row != expectedRow {
			t.Errorf(errMsg, expectedRow, cursor.Row)
		}
	}
}

// Tests the move up function of the cursor,
// when the cursor is at the top border (row 0).
func TestMoveUpAtBorder(t *testing.T) {
	errMsg := "Cursor.MoveUp: "
	errMsg += "Expected row to not change when moving up, "
	errMsg += "but got %d."

	cursor := NewCursor()
	cursor.MoveUp()

	if cursor.Row != 0 {
		t.Errorf(errMsg, cursor.Row)
	}
}

// Tests the move down function of the cursor.
func TestMoveDown(t *testing.T) {
	errMsg := "Cursor.MoveDown: "
	errMsg += "Expected row to be %d after moving down, "
	errMsg += "but got %d."

	cursor := NewCursor()
	expectedRow := 0

	for i := 0; i < 8; i++ {
		expectedRow++
		cursor.MoveDown()

		if cursor.Row != expectedRow {
			t.Errorf(errMsg, expectedRow, cursor.Row)
		}
	}
}

// Tests the move down function of the cursor,
// when the cursor is at the bottom border (row 8).
func TestMoveDownAtBorder(t *testing.T) {
	errMsg := "Cursor.MoveDown: "
	errMsg += "Expected row to not change when moving down, "
	errMsg += "but got %d."

	cursor := NewCursor()
	cursor.Row = 8
	cursor.MoveDown()

	if cursor.Row != 8 {
		t.Errorf(errMsg, cursor.Row)
	}
}
