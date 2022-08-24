package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"net/http"
)

func GetCommentById(key string) (int, map[string]interface{}) {
	var comments models.Comment
	result := config.DB.Preload("Customer.User.Role").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&comments, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func GetCommentsByStationeryId(key string) (int, map[string]interface{}) {
	var comments []models.Comment
	result := config.DB.Preload("Customer.User.Role").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&comments, "stationery_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func GetCommentsByCustomerId(key string) (int, map[string]interface{}) {
	var comments []models.Comment
	result := config.DB.Preload("Customer.User.Role").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&comments, "customer_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func AddComment(comment models.Comment) (int, map[string]interface{}) {
	result := config.DB.Create(&comment)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = comment
		return http.StatusCreated, send
	}
}

func UpdateComment(comment models.Comment, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&comment)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = comment
		return http.StatusCreated, send
	}
}
