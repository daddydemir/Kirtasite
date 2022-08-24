package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"net/http"
)

func GetFilesById(key string) (int, map[string]interface{}) {
	var files models.File
	result := config.DB.Preload("Customer.User.Role").Find(&files, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = files
		return http.StatusOK, send
	}
}

func GetFilesByCustomerId(key string) (int, map[string]interface{}) {
	var files []models.File
	result := config.DB.Preload("Customer.User.Role").Find(&files, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = files
		return http.StatusOK, send
	}
}

func AddFiles(file models.File) (int, map[string]interface{}) {
	result := config.DB.Create(&file)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = file
		return http.StatusCreated, send
	}
}

func UpdateFiles(file models.File, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&file)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = file
		return http.StatusCreated, send
	}
}
