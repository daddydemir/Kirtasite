package cache

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"Kirtasite/pkg/db"
	"context"
	"log"
	"time"
)

func ReadToken(sessionId, token string) bool {
	check := config.RDB.Ping(context.Background()).Err()
	if check != nil {
		log.Println("Redis server is unreachable.")
		return db.CheckSession(sessionId, token)
	}
	val := config.RDB.Get(context.Background(), sessionId)
	if val.Err() != nil {
		log.Println("key doesn't exist.")
		log.Println(val.Err())
		return false
	}
	if val.Val() == token {
		log.Println("session is truth.")
		return true
		// problem - only one session can be opened at a time!
	} else {
		log.Println("Redis has returned: ", val.Val(), val.Err())
		return false
	}
}

func WriteToken(session models.Session) bool {
	log.Println("TIME", session.ExDate.Sub(time.Now()))
	resp := config.RDB.Set(context.Background(), session.Id.String(), session.Token, session.ExDate.Sub(time.Now()))
	if resp.Err() != nil {
		log.Fatal(resp.Err())
		return false
	} else {
		return true
	}
}
