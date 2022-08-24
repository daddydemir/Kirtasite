package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"net/http"
)

func GetPrices() (int, map[string]interface{}) {
	var prices []models.Price
	return all(prices)
}

func GetPriceById(key string) (int, map[string]interface{}) {
	var price models.Price
	return query("id = ? ", key, &price)
}

func GetPriceByStationeryId(key string) (int, map[string]interface{}) {
	var prices []models.Price
	return query("stationery_id = ?", key, prices)
}

func AddPrice(price models.Price) (int, map[string]interface{}) {
	result := config.DB.Create(&price)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = price
		return http.StatusCreated, send
	}
}

func UpdatePrice(prices models.Price, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&prices)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = prices
		return http.StatusCreated, send
	}
}
