package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattladany/redis-manager/internal/tui"
)

func initialModel() tui.Model {
	return tui.Model{
		SortedKeys: make([]string, 0),
		KeyValues:  make(map[string]string),
		Selected:   make(map[int]struct{}),
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
