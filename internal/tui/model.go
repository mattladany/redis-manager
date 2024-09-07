package tui

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type State int

const (
	Error State = iota
	Active
	Fetching
)

type Model struct {
	List  list.Model
	State State
}

type item struct {
	key, value string
}

func (i item) Title() string       { return i.key }
func (i item) Description() string { return i.value }
func (i item) FilterValue() string { return i.key }

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		return tea.KeyMsg{Type: tea.KeyType(tea.KeyCtrlR)}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	docStyle := lipgloss.NewStyle().Margin(1, 2)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)

	// Key press
	case tea.KeyMsg:
		switch msg.String() {

		// Exit
		case "ctrl+c", "q":
			// Clear the screen
			fmt.Print("\033[2J")
			fmt.Print("\033[H")
			return m, tea.Quit

		// HOT KEYS

		// Refresh Redis data
		case "ctrl+r":
			m.State = Fetching
			return m, FetchAll()
		}

	// Update Redis data
	case FetchResponse:
		m.State = Active
		items := make([]list.Item, 0)
		for key, value := range msg.Data {
			items = append(items, item{key: key, value: value})
		}

		sort.Slice(items, func(i, j int) bool {
			return items[i].FilterValue() < items[j].FilterValue()
		})

		m.List.SetItems(items)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return render(m)
}
