package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func GetAllData() (map[string]string, error) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:30000",
		Password: "",
		DB:       0,
	})

	keyValues := make(map[string]string)
	iter := rdb.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		value, err := rdb.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		keyValues[key] = value
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	return keyValues, nil
}
