package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"net/http"
	"strconv"
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

func AddPrice(price models.Price, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	var user models.User
	config.DB.Preload("Role").Find(&user, "id = ?", userId)
	if user.Role.Name != core.STATIONERY {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	result := config.DB.Create(&price)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Created)
		send["data"] = price
		return http.StatusCreated, send
	}
}

func UpdatePrice(prices models.Price, key string, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)

	var price models.Price
	config.DB.Find(&price, "stationery_id = ?", userId)

	if userId != strconv.Itoa(price.StationeryId) {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	prices.Id, _ = strconv.Atoi(key)
	result := config.DB.Model(&models.Price{}).Where("id = ?", key).Save(&prices)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Updated)
		send["data"] = prices
		return http.StatusCreated, send
	}
}
