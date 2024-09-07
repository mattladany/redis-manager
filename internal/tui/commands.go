package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattladany/redis-manager/internal/redis"
)

type FetchResponse struct {
	Data map[string]string
	Err  error
}

func FetchAll() tea.Cmd {
	return func() tea.Msg {
		data, err := redis.GetAllData()
		return FetchResponse{
			Data: data,
			Err:  err,
		}
	}
}
