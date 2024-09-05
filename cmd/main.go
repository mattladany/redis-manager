package main

import (
	"fmt"
	"os"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattladany/redis-manager/internal/redis"
	"github.com/mattladany/redis-manager/internal/tui"
)

func initialModel() tui.Model {

	data := redis.GetAllData()
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return tui.Model{
		SortedKeys: keys,
		KeyValues:  data,
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
