package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"Kirtasite/secuirty"
	"net/http"
	"strconv"
)

func GetStationeries() (int, map[string]interface{}) {
	var stationeries []models.Stationery
	result := config.DB.Preload("User.Role").Preload("Address.City").Preload("Address.District").Find(&stationeries)
	if result.Error != nil {
		return http.StatusNoContent, core.SendMessage(core.NoContent)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = stationeries
		return http.StatusOK, send
	}
}

func GetStationeryById(key string) (int, map[string]interface{}) {
	var stationery models.Stationery
	result := config.DB.Preload("User.Role").Preload("Address.City").Preload("Address.District").Find(&stationery, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, core.SendMessage(core.NoContent)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = stationery
		return http.StatusOK, send
	}
}

func AddStationery(stationery models.Stationery) (int, map[string]interface{}) {
	stationery.User.Password = secuirty.HashPassword(stationery.User.Password)
	result := config.DB.Create(&stationery)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Created)
		send["data"] = stationery
		return http.StatusCreated, send
	}
}

func UpdateStationery(stationery models.Stationery, key string, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	if userId != key {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	var tempUser models.User
	config.DB.Find(&tempUser, "id = ?", key)
	stationery.User = tempUser
	stationery.UserId, _ = strconv.Atoi(key)
	result := config.DB.Model(&models.Stationery{}).Where("user_id = ?", key).Save(&stationery)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Updated)
		send["data"] = stationery
		return http.StatusCreated, send
	}
}
