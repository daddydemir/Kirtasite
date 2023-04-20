package main

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"Kirtasite/pkg/cache"
	"github.com/google/uuid"
)

func main() {
	config.RedisClient()
	//cache.ReadToken(models.User{}, "token")
	//cache.WriteToken("")
	var m models.Session
	m.Id = uuid.New()
	//log.Println(m.Id)

	cache.ReadToken("name1", "mehmet")
}
