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
)

func headerWindow(width int) string {
	return windowStyle.Width(width).Render(titleStyle.Render(HeaderWindowTitle))
}

func leftWindow(keys []string, cursorIndex int, width int) string {
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

func rightWindow(key string, value string, width int) string {
	content := titleStyle.Render("Selected Key Value") + "\n\n"
	content += fmt.Sprintf("Key: %s\nValue: %s", key, value)
	return windowStyle.Width(width).Render(content)
}

func bottomWindow(content string, width int) string {
	return windowStyle.Width(width).Render(content)
}

func render(m Model) string {
	// Calculate the width for each panel (assuming a total width of 100)
	totalWidth := 100
	leftWidth := totalWidth / 2
	rightWidth := totalWidth - leftWidth

	// Create top window
	topWindow := headerWindow(totalWidth)

	// Create middle windows
	leftWindow := leftWindow(m.SortedKeys, m.Cursor, leftWidth)
	rightWindow := rightWindow(m.SortedKeys[m.Cursor], m.KeyValues[m.SortedKeys[m.Cursor]], rightWidth)
	middleContent := lipgloss.JoinHorizontal(lipgloss.Top, leftWindow, rightWindow)

	// Create bottom window
	bottomWindow := bottomWindow("Press 'q' to quit", totalWidth)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		topWindow,
		middleContent,
		bottomWindow,
	)
}
