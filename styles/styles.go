
package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	PrimaryColor   = lipgloss.Color("#FF4B6C") // Vibrant pink
	SecondaryColor = lipgloss.Color("#4ECDC4") // Teal accent
	textColor      = lipgloss.Color("#F7F7F7") // Soft white
	categoryColor  = lipgloss.Color("#FFB86C") // Warm orange
	accentColor    = lipgloss.Color("#BD93F9") // Purple accent

	// Styles
	LogoStyle = lipgloss.NewStyle().
		Foreground(PrimaryColor).
		Bold(true).
		Padding(2, 4).
		Align(lipgloss.Center).
		Border(lipgloss.DoubleBorder()).
		BorderForeground(SecondaryColor).
		Width(70)

	CategoryStyle = lipgloss.NewStyle().
		Foreground(categoryColor).
		Bold(true).
		Padding(1, 0, 0, 4).
		Underline(true)

	ProductStyle = lipgloss.NewStyle().
		Foreground(textColor).
		Padding(0, 4).
		MarginLeft(2)

	PriceStyle = lipgloss.NewStyle().
		Foreground(accentColor).
		Bold(true).
		Width(10).
		Align(lipgloss.Right)

	StarStyle = lipgloss.NewStyle().
		Foreground(SecondaryColor).
		Bold(true).
		Padding(0, 0, 0, 1)

	SeparatorStyle = lipgloss.NewStyle().
		Foreground(categoryColor).
		Padding(0, 2)
)
