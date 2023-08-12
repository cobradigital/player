package redis

import (
	"context"
	"os"
	"repo/configs"
	"repo/loggers"
	"time"

	"github.com/redis/go-redis/v9"
)

var log = loggers.Get()

var ctx = context.Background()

var Client *redis.Client

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.MustGetString("redis.host"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	Client = rdb

	res := Client.Ping(ctx)

	if res.String() != "ping: PONG" {
		log.Info("Failed Connecting Redis DB. err : ", res)
		os.Exit(2)
	}

	log.Info("Success Connecting Redis DB.")
}

func Set(key, value string, t time.Duration) error {
	return Client.Set(ctx, key, value, t).Err()
}

func Get(key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

func Del(key string) (int64, error) {
	return Client.Del(ctx, key).Result()
}
