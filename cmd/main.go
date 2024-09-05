package main

import (
	"context"
	"fmt"
	"os"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/redis/go-redis/v9"
)

type model struct {
	keys      []string
	keyValues map[string]string
	cursor    int
	selected  map[int]struct{}
}

var (
	windowStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("33")).
			Bold(true)
)

func initialModel() model {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:30000",
		Password: "",
		DB:       0,
	})

	keyValues := make(map[string]string)
	keys := make([]string, 0)
	iter := rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		value, err := rdb.Get(ctx, key).Result()
		if err != nil {
			panic(err)
		}
		keys = append(keys, key)
		keyValues[key] = value
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	sort.Strings(keys)

	return model{
		keys:      keys,
		keyValues: keyValues,
		selected:  make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.keyValues)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// Calculate the width for each panel (assuming a total width of 100)
	totalWidth := 100
	leftWidth := totalWidth / 2
	rightWidth := totalWidth - leftWidth

	// Left panel content
	leftContent := titleStyle.Render("Redis Keys") + "\n\n"
	for i, key := range m.keys {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		leftContent += fmt.Sprintf("%s %s\n", cursor, key)
	}

	// Right panel content
	rightContent := titleStyle.Render("Selected Key Value") + "\n\n"
	if m.cursor < len(m.keys) {
		selectedKey := m.keys[m.cursor]
		rightContent += fmt.Sprintf("Key: %s\nValue: %s", selectedKey, m.keyValues[selectedKey])
	}

	// Create windows
	leftWindow := windowStyle.Width(leftWidth).Render(leftContent)
	rightWindow := windowStyle.Width(rightWidth).Render(rightContent)

	// Create top window
	topWindow := windowStyle.Width(totalWidth).Render(titleStyle.Render("Redis Manager"))

	// Create bottom window
	bottomWindow := windowStyle.Width(totalWidth).Render("Press 'q' to quit")

	// Combine windows
	middleContent := lipgloss.JoinHorizontal(lipgloss.Top, leftWindow, rightWindow)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		topWindow,
		middleContent,
		bottomWindow,
	)

}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
