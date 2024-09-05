package main

import (
	"context"
	"fmt"
	"os"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattladany/redis-manager/internal/tui"
	"github.com/redis/go-redis/v9"
)

func initialModel() tui.Model {

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

	return tui.Model{
		Keys:      keys,
		KeyValues: keyValues,
		Selected:  make(map[int]struct{}),
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
