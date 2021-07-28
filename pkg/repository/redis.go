package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

var ctx = context.Background()

func NewRedisDB(cfg Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return rdb, err
	}
	return rdb, nil
}
