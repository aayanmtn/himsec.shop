package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	primaryColor = lipgloss.Color("#D4A373")
	textColor    = lipgloss.Color("#463F3A")

	// Styles
	TitleStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true).
		Padding(1, 0, 1, 2)

	ProductStyle = lipgloss.NewStyle().
		Foreground(textColor).
		Padding(0, 2)

	SelectedStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true).
		Padding(0, 2)

	InfoStyle = lipgloss.NewStyle().
		Foreground(textColor).
		Italic(true).
		Padding(1, 2)
)