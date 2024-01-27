package config

import (
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
)

var once sync.Once
var RDB *redis.Client

func RedisClient() {
	RDB = getInstance()
}

func getInstance() *redis.Client {

	once.Do(func() {
		db, _ := strconv.Atoi(Get("REDIS_DB"))
		RDB = redis.NewClient(&redis.Options{
			Addr:     Get("REDIS_ADDR"),
			Password: Get("REDIS_PASS"),
			DB:       db,
		})
	})

	return RDB
}
