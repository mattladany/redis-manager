package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattladany/redis-manager/internal/redis"
)

type RefreshDataCmd struct {
	Data map[string]string
	Err  error
}

func RefreshData() tea.Cmd {
	return func() tea.Msg {
		data, err := redis.GetAllData()
		return RefreshDataCmd{
			Data: data,
			Err:  err,
		}
	}
}
