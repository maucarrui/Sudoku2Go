package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var selectedValueStyle = lipgloss.NewStyle().
	Bold(true).
	Render

var initialValueStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#c4bebe")).
	Render

var placedValuesStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#59a6d9")).
	Render

var matchingValuesStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#92d959")).
	Render

var conflictValueStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#92000a")).
	Render

var sudokuStyle = lipgloss.NewStyle().
	Width(60).
	Padding(5).
	Render

var instructionsStyle = lipgloss.NewStyle().
	Width(40).
	Padding(5).
	Render
