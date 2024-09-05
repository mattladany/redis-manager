package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	windowStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("33")).
			Bold(true)
)

const (
	HeaderWindowTitle = "Redis Manager"
	FooterWindowTitle = "'q' quit, '^r' refresh"
)

func formatHeaderWindow(width int) string {
	return windowStyle.Width(width).Render(titleStyle.Render(HeaderWindowTitle))
}

func formatLeftWindow(keys []string, cursorIndex int, width int) string {
	content := titleStyle.Render("Redis Keys") + "\n\n"
	for i, key := range keys {
		cursor := " "
		if cursorIndex == i {
			cursor = ">"
		}
		content += fmt.Sprintf("%s %s\n", cursor, key)
	}
	return windowStyle.Width(width).Render(content)
}

func formatRightWindow(key string, value string, width int) string {
	content := titleStyle.Render("Selected Key Value") + "\n\n"
	content += fmt.Sprintf("Key: %s\nValue: %s", key, value)
	return windowStyle.Width(width).Render(content)
}

func formatBottomWindow(width int) string {
	return windowStyle.Width(width).Render(FooterWindowTitle)
}

func render(m Model) string {
	switch m.State {
	case Fetching:
		return renderFetching()
	case Active:
		return renderActive(m)
	default:
		return "Unknown state"
	}
}

func renderFetching() string {
	return "Fetching data..."
}

func renderActive(m Model) string {

	// Calculate the width for each panel (assuming a total width of 100)
	totalWidth := 100
	leftWidth := totalWidth / 2
	rightWidth := totalWidth - leftWidth

	// Create top window
	topWindow := formatHeaderWindow(totalWidth)

	// Create middle windows
	leftWindow := "No data found"
	rightWindow := "No data found"
	if len(m.SortedKeys) > 0 {
		leftWindow = formatLeftWindow(m.SortedKeys, m.Cursor, leftWidth)
		rightWindow = formatRightWindow(m.SortedKeys[m.Cursor], m.KeyValues[m.SortedKeys[m.Cursor]], rightWidth)
	}
	middleContent := lipgloss.JoinHorizontal(lipgloss.Top, leftWindow, rightWindow)

	// Create bottom window
	bottomWindow := formatBottomWindow(totalWidth)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		topWindow,
		middleContent,
		bottomWindow,
	)
}
