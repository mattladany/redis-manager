package tui

import (
	"fmt"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	Error State = iota
	Active
	Fetching
)

type Model struct {
	SortedKeys []string
	KeyValues  map[string]string
	Cursor     int
	Selected   map[int]struct{}
	State      State
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		return tea.KeyMsg{Type: tea.KeyType(tea.KeyCtrlR)}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			// Clear the screen
			fmt.Print("\033[2J")
			fmt.Print("\033[H")
			return m, tea.Quit

		// NAVIGATION

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Cursor < len(m.KeyValues)-1 {
				m.Cursor++
			}

		// HOT KEYS

		// Refresh the data
		case "ctrl+r":
			m.State = Fetching
			return m, RefreshData()
		}

	// Update the data and reset the cursor to 0
	case RefreshDataCmd:
		m.State = Active
		m.KeyValues = msg.Data
		m.SortedKeys = make([]string, 0, len(m.KeyValues))
		for key := range m.KeyValues {
			m.SortedKeys = append(m.SortedKeys, key)
		}
		sort.Strings(m.SortedKeys)
		m.Cursor = 0
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m Model) View() string {
	return render(m)
}
