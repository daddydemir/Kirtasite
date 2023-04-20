package db

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"log"
)

func CheckSession(sessionId, token string) bool {
	//todo: check token expire date!
	var session models.Session
	config.DB.Find(&session, "id = ?", sessionId)
	if token == session.Token {
		log.Println("token is invalid.")
		return true
	}
	return false
}
