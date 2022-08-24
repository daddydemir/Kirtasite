package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"net/http"
)

func GetCities() (int, map[string]interface{}) {
	var cities []models.City
	return all(cities)
}

func GetCityById(key string) (int, map[string]interface{}) {
	var city models.City
	result := config.DB.Find(&city, "id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = city
		return http.StatusOK, send
	}
}
