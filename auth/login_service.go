package auth

import (
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"Kirtasite/pkg/cache"
	"Kirtasite/secuirty"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"strconv"
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
		message["token"] = GenerateToken(strconv.Itoa(user.Id))
		res := createSession(user, GenerateToken(strconv.Itoa(user.Id)), r)
		s := cache.WriteToken(res)
		log.Println(s)
		return http.StatusOK, message
	}
	return http.StatusBadRequest, core.SendMessage(core.LoginFail)
}

func createSession(user models.User, token string, r *http.Request) models.Session {
	var session models.Session
	session.Email = user.Mail
	session.Ip = r.RemoteAddr
	session.Token = token
	session.CrDate = time.Now()
	session.ExDate = time.Now().Add(time.Hour * 10)
	config.DB.Clauses(clause.Returning{}).Create(&session)
	log.Println(session.Id)
	return session
}
