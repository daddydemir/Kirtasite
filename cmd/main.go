package main

import (
	"fmt"
	"time"
)

func main() {
	//config.RedisClient()
	//cache.ReadToken(models.User{}, "token")
	//cache.WriteToken("")
	t := time.Now() // Time
	z := time.Now().Add(-time.Hour * 10)
	fmt.Println(t.Sub(z))

	fmt.Println(time.Now().Sub(time.Now().Add(-time.Hour * 10)))
	fmt.Println(time.Hour * 10)
}
