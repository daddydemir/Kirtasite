package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetStationeries() (int, map[string]interface{}) {
	var stationeries []models.Stationeries
	result := config.DB.Preload(clause.Associations).Find(&stationeries)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = stationeries
		return http.StatusOK, send
	}
}

func GetStationeryById(key string) (int, map[string]interface{}) {
	var stationery models.Stationeries
	result := config.DB.Preload(clause.Associations).Find(&stationery, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = stationery
		return http.StatusOK, send
	}
}

func AddStationery(stationery models.Stationeries) (int, map[string]interface{}) {
	result := config.DB.Create(&stationery)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = stationery
		return http.StatusCreated, send
	}
}

func UpdateStationery(stationery models.Stationeries, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&stationery)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = stationery
		return http.StatusCreated, send
	}
}
