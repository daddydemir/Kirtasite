package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"net/http"
)

func GetPrices() (int, map[string]interface{}) {
	var prices []models.Prices
	return all(prices)
}

func GetPriceById(key string) (int, map[string]interface{}) {
	var price models.Prices
	return query("id = ? ", key, &price)
}

func GetPriceByStationeryId(key string) (int, map[string]interface{}) {
	var prices []models.Prices
	return query("stationery_id = ?", key, prices)
}

func AddPrice(price models.Prices) (int, map[string]interface{}) {
	result := config.DB.Create(&price)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = price
		return http.StatusCreated, send
	}
}

func UpdatePrice(prices models.Prices, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&prices)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = prices
		return http.StatusCreated, send
	}
}