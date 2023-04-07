package cache

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"context"
	"fmt"
	"log"
	"time"
)

func ReadToken(token string) {
	val := config.RDB.Get(context.Background(), "auth.TokenParser(token)")
	if val.Err() != nil {
		log.Println("key doesn't exist.")
	}
	if val.Val() == token {
		log.Println("session is truth.")
		// problem - only one session can be opened at a time!
	}
	fmt.Println("Redis has returned: ", val.Val(), val.Err())
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
