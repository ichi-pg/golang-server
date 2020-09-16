package redis

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/ichi-pg/golang-server/internal/pkg/env"
)

func runWithClient(f func(*redis.Client) error) error {
	cli := redis.NewClient(&redis.Options{
		Addr: os.Getenv(env.RedisAddr),
	})
	defer cli.Close()
	return f(cli)
}
