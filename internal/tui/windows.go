package tui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	HeaderWindowTitle = "Redis Manager"
	FooterWindowTitle = "q: quit • ^r refresh"
)

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
	return lipgloss.NewStyle().Margin(1, 2).Render(m.List.View())

}
