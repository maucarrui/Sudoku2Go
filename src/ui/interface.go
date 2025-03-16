package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maucarrui/Sudoku2Go/internal"
)

func (game Game) Init() tea.Cmd {
	return game.timer.Init()
}

func (game Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Define the keys to exit the game.
	exitKeys := internal.NewStringSetFromValues(
		[]string{"ctrl+c", "q", "esc"},
	)

	// Define the keys to move the cursor.
	movementKeys := internal.NewStringSetFromValues(
		[]string{
			"ctrl+p", "ctrl+n", "ctrl+b", "ctrl+f",
			"up", "down", "left", "right",
			"w", "s", "a", "d",
		},
	)

	// Define the keys to insert values into the sudoku.
	numberKeys := internal.NewStringSetFromValues(
		[]string{
			"1", "2", "3", "4", "5", "6", "7", "8", "9",
		},
	)

	// Define the keys to remove values.
	deleteKeys := internal.NewStringSetFromValues(
		[]string{
			"backspace",
		},
	)

	// Define the keys to enable/disable blinking cursor.
	blinkingKeys := internal.NewStringSetFromValues(
		[]string{
			"t",
		},
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		if exitKeys.Contains(key) {
			// Exit game.
			game.message = "Exiting Game..."
			return game, tea.Quit
		} else if movementKeys.Contains(key) {
			// Move the cursor.
			game.MoveCursor(key)
		} else if numberKeys.Contains(key) {
			// Insert a value in the Sudoku.
			game.AddValue(key)
		} else if deleteKeys.Contains(key) {
			// Delete the value where the cursor is currently on.
			game.RemoveValue()
		} else if blinkingKeys.Contains(key) {
			game.blinking = !game.blinking
		} else {
			// Default scenario.
			game.message = "Not entering any scenario for key: " + key
		}

		// Check if game has been completed.
		if game.IsComplete() {
			// Move the cursor out of bounds to show the complete sudoku.
			game.cursor.Row = -1
			game.cursor.Col = -1

			return game, tea.Quit
		}
	}

	var cmd tea.Cmd
	game.timer, cmd = game.timer.Update(msg)

	return game, cmd
}

func (game Game) View() string {
	sudoku := gameToString(game)
	sudoku += "Elapsed Time: " + PrintElapsedTime(game.timer.Elapsed()) + "\n"

	sudoku = sudokuStyle(sudoku)

	instructions := instructionsStyle(
		RulesToString() + "\n\n" + ControlsToString(game),
	)

	return lipgloss.JoinHorizontal(lipgloss.Center, instructions, sudoku)
}
