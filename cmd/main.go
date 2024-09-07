package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattladany/redis-manager/internal/tui"
)

const (
	HeaderWindowTitle = "Redis Manager"
)

func initialModel() tui.Model {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = HeaderWindowTitle
	return tui.Model{
		List: list,
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
