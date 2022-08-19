package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

func GetAddressById(key string) (int, map[string]interface{}) {
	var address models.Addresses
	result := config.DB.Preload(clause.Associations).Find(&address, "id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = address
		return http.StatusOK, send
	}
}

func GetAddressByCityId(key string) (int, map[string]interface{}) {
	var addresses []models.Addresses
	result := config.DB.Preload(clause.Associations).Find(&addresses, "city_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		msg := SendMessage(Ok)
		msg["data"] = addresses
		return http.StatusOK, msg
	}
}

func AddAddress(address models.Addresses) (int, map[string]interface{}) {
	result := config.DB.Create(&address)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		msg := SendMessage(Created)
		msg["data"] = address
		return http.StatusCreated, msg
	}
}

func UpdateAddress(addresses models.Addresses, key string) (int, map[string]interface{}) {
	// TODO : auth eksik
	// delete this (if)
	id, _ := strconv.Atoi(key)
	if id != addresses.Id {
		return http.StatusForbidden, SendMessage(Forbidden)
	}
	return save(addresses)
}
