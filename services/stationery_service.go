package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"Kirtasite/secuirty"
	"fmt"
	"net/http"
)

func GetStationeries() (int, map[string]interface{}) {
	var stationeries []models.Stationery
	result := config.DB.Preload("User.Role").Preload("Address.City").Preload("Address.District").Find(&stationeries)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = stationeries
		return http.StatusOK, send
	}
}

func GetStationeryById(key string) (int, map[string]interface{}) {
	var stationery models.Stationery
	result := config.DB.Preload("User.Role").Preload("Address.City").Preload("Address.District").Find(&stationery, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = stationery
		return http.StatusOK, send
	}
}

func AddStationery(stationery models.Stationery) (int, map[string]interface{}) {
	stationery.User.Password = secuirty.HashPassword(stationery.User.Password)
	result := config.DB.Create(&stationery)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = stationery
		return http.StatusCreated, send
	}
}

func UpdateStationery(stationery models.Stationery, key string) (int, map[string]interface{}) {
	fmt.Println(stationery)
	result := config.DB.Model(&models.Stationery{}).Where("user_id = ?", key).Save(&stationery)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = stationery
		return http.StatusCreated, send
	}
}
