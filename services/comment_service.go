package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetCommentById(key string) (int, map[string]interface{}) {
	var comments models.Comments
	result := config.DB.Preload(clause.Associations).Find(&comments, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func GetCommentsByStationeryId(key string) (int, map[string]interface{}) {
	var comments []models.Comments
	result := config.DB.Preload(clause.Associations).Find(&comments, "stationery_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func GetCommentsByCustomerId(key string) (int, map[string]interface{}) {
	var comments []models.Comments
	result := config.DB.Preload(clause.Associations).Find(&comments, "customer_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func AddComment(comment models.Comments) (int, map[string]interface{}) {
	result := config.DB.Create(&comment)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = comment
		return http.StatusCreated, send
	}
}

func UpdateComment(comment models.Comments, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&comment)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = comment
		return http.StatusCreated, send
	}
}
