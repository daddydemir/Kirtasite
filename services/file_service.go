package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"net/http"
	"reflect"
	"strconv"
)

func GetFilesById(key string, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	var user models.User
	config.DB.Preload("Role").Find(&user, "id = ?", userId)
	var myfile models.File
	config.DB.Find(&myfile, "id = ?", key)
	if reflect.DeepEqual(userId, string(rune(myfile.CustomerId))) {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}

	var files models.File
	result := config.DB.Preload("Customer.User.Role").Find(&files, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = files
		return http.StatusOK, send
	}
}

func GetFilesByCustomerId(key string, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	var myfile models.File
	config.DB.Find(&myfile, "id = ?", key)
	if userId != string(myfile.CustomerId) {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}

	var files []models.File
	result := config.DB.Preload("Customer.User.Role").Find(&files, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = files
		return http.StatusOK, send
	}
}

func AddFiles(file models.File, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	var customer models.Customer
	config.DB.Preload("User.Role").Find(&customer, "user_id = ?", userId)
	if customer.User.Role.Name != core.CUSTOMER {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	result := config.DB.Create(&file)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Created)
		send["data"] = file
		return http.StatusCreated, send
	}
}

func UpdateFiles(file models.File, key string, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	var myfile models.File
	config.DB.Find(&myfile, "id = ?", key)
	if userId != string(myfile.CustomerId) {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	file.Id, _ = strconv.Atoi(key)

	result := config.DB.Model(&models.File{}).Where("id = ?", key).Save(&file)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Updated)
		send["data"] = file
		return http.StatusCreated, send
	}
}
