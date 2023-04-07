package main

import (
	"Kirtasite/models"
	"github.com/google/uuid"
)

func main() {
	//config.RedisClient()
	//cache.ReadToken(models.User{}, "token")
	//cache.WriteToken("")
	var m models.Session
	m.Id = uuid.New()
	println(m.Id)
}
