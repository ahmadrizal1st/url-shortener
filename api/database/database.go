package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbNo,
	})

	return rdb
}

func Inizialize() {
	Client = CreateClient(0)
}
