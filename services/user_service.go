package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"Kirtasite/secuirty"
	"net/http"
)

func ImageUpload(url string, key string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	if userId != key {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	var user models.User
	_ = config.DB.Find(&user, "id = ?", key)
	user.ImagePath = url
	result := config.DB.Model(&models.User{}).Where("id = ?", key).Save(&user)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	send := core.SendMessage(core.Updated)
	send["data"] = user
	return http.StatusOK, send
}

func PasswordUpdate(lastpass string, againpass string, newpass string, key string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	if userId != key {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	if lastpass != againpass {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	var user models.User
	config.DB.Find(&user, "id = ?", userId)
	isTure := secuirty.CheckPassword(user.Password, lastpass)
	if !isTure {
		return http.StatusBadRequest, core.SendMessage(core.NoPass)
	}
	user.Password = secuirty.HashPassword(newpass)
	result := config.DB.Model(&models.User{}).Where("id = ?", key).Save(&user)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = user
		return http.StatusOK, send
	}
}
