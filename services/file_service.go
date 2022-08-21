package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetFilesById(key string) (int, map[string]interface{}) {
	var files models.Files
	result := config.DB.Preload(clause.Associations).Find(&files, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = files
		return http.StatusOK, send
	}
}

func GetFilesByCustomerId(key string) (int, map[string]interface{}) {
	var files []models.Files
	result := config.DB.Preload(clause.Associations).Find(&files, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = files
		return http.StatusOK, send
	}
}

func AddFiles(file models.Files) (int, map[string]interface{}) {
	result := config.DB.Create(&file)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = file
		return http.StatusCreated, send
	}
}

func UpdateFiles(file models.Files, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&file)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = file
		return http.StatusCreated, send
	}
}
