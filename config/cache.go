package config

import (
	"github.com/redis/go-redis/v9"
	"strconv"
)

var RDB *redis.Client

func RedisClient() {
	db, _ := strconv.Atoi(Get("REDIS_DB"))
	RDB = redis.NewClient(&redis.Options{
		Addr:     Get("REDIS_ADDR"),
		Password: Get("REDIS_PASS"),
		DB:       db,
	})
}
