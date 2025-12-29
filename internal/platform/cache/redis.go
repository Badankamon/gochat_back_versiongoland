package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/Badankamon/gochat_backend/internal/config"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Connect(cfg config.RedisConfig) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := RDB.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")
}
