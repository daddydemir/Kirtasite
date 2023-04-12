package auth

import (
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"Kirtasite/secuirty"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func LoginService(mail, password string, r *http.Request) (int, interface{}) {
	var user models.User
	config.DB.Find(&user, "mail = ?", mail)
	if user.Id == 0 {
		return http.StatusBadRequest, core.SendMessage(core.LoginFail)
	}
	isTrue := secuirty.CheckPassword(user.Password, password)
	if isTrue {
		message := core.SendMessage(core.Ok)
		message["data"] = user
		message["token"] = GenerateToken(createSession(user, r))
		return http.StatusOK, message
	}
	return http.StatusBadRequest, core.SendMessage(core.LoginFail)
}

func createSession(user models.User, r *http.Request) models.Session {
	var session models.Session
	session.Id = uuid.New()
	session.Email = user.Mail
	session.Ip = r.RemoteAddr
	session.CrDate = time.Now()
	session.ExDate = time.Now().Add(time.Hour * 10)
	log.Println(session.Id)
	return session
}
