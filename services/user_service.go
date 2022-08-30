package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
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
