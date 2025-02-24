package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	PrimaryColor   = lipgloss.Color("#FF6B35") // Orange color for brand
	textColor      = lipgloss.Color("#FFFFFF") // White text
	categoryColor  = lipgloss.Color("#FF8B5B") // Lighter orange for categories

	// Styles
	LogoStyle = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Bold(true).
		Padding(2, 0).
		Align(lipgloss.Center).
		Border(lipgloss.NormalBorder()).
		BorderForeground(PrimaryColor).
		Width(60) // Increased width to make logo more prominent

	CategoryStyle = lipgloss.NewStyle().
		Foreground(categoryColor).
		Bold(true).
		Padding(1, 0, 0, 2)

	ProductStyle = lipgloss.NewStyle().
		Foreground(textColor).
		Padding(0, 2)

	PriceStyle = lipgloss.NewStyle().
		Foreground(textColor).
		Width(8).
		Align(lipgloss.Right)

	StarStyle = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Padding(0, 0, 0, 1)

	SeparatorStyle = lipgloss.NewStyle().
		Foreground(categoryColor).
		Padding(0, 2)
)