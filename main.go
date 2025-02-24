package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/maucarrui/Sudoku2Go/src/ui"
	"os"
)

func main() {
	program := tea.NewProgram(ui.NewGame(), tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Printf("An error occurred during execution: %v", err)
		os.Exit(1)
	}
}
