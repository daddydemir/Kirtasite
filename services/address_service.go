package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"net/http"
	"strconv"
)

func GetAddressById(key string) (int, map[string]interface{}) {
	var address models.Address
	result := config.DB.Preload("City").Preload("District").Find(&address, "id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, core.SendMessage(core.NoContent)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = address
		return http.StatusOK, send
	}
}

func GetAddressByCityId(key string) (int, map[string]interface{}) {
	var addresses []models.Address
	result := config.DB.Preload("City").Preload("District").Find(&addresses, "city_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, core.SendMessage(core.NoContent)
	} else {
		msg := core.SendMessage(core.Ok)
		msg["data"] = addresses
		return http.StatusOK, msg
	}
}

func AddAddress(address models.Address) (int, map[string]interface{}) {
	result := config.DB.Create(&address)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		msg := core.SendMessage(core.Created)
		msg["data"] = address
		return http.StatusCreated, msg
	}
}

func UpdateAddress(addresses models.Address, key string, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	var stationery models.Stationery
	config.DB.Find(&stationery, "user_id = ?", userId)
	address_id, _ := strconv.Atoi(key)
	if stationery.AddressId != address_id {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	addresses.Id = address_id
	result := config.DB.Model(&models.Address{}).Where("id = ?", key).Save(&addresses)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		msg := core.SendMessage(core.Updated)
		msg["data"] = addresses
		return http.StatusCreated, msg
	}
}
