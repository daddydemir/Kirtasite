package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"net/http"
	"strconv"
)

func GetAddressById(key string) (int, map[string]interface{}) {
	var address models.Address
	result := config.DB.Preload("City").Preload("District").Find(&address, "id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = address
		return http.StatusOK, send
	}
}

func GetAddressByCityId(key string) (int, map[string]interface{}) {
	var addresses []models.Address
	result := config.DB.Preload("City").Preload("District").Find(&addresses, "city_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		msg := SendMessage(Ok)
		msg["data"] = addresses
		return http.StatusOK, msg
	}
}

func AddAddress(address models.Address) (int, map[string]interface{}) {
	result := config.DB.Create(&address)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		msg := SendMessage(Created)
		msg["data"] = address
		return http.StatusCreated, msg
	}
}

func UpdateAddress(addresses models.Address, key string) (int, map[string]interface{}) {
	// TODO : auth eksik
	// delete this (if)
	id, _ := strconv.Atoi(key)
	if id != addresses.Id {
		return http.StatusForbidden, SendMessage(Forbidden)
	}
	return save(addresses)
}
