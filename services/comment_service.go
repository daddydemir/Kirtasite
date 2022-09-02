package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"net/http"
	"strconv"
)

func GetCommentById(key string) (int, map[string]interface{}) {
	var comments models.Comment
	result := config.DB.Preload("Customer.User.Role").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&comments, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func GetCommentsByStationeryId(key string) (int, map[string]interface{}) {
	var comments []models.Comment
	result := config.DB.Preload("Customer.User.Role").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&comments, "stationery_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func GetCommentsByCustomerId(key string) (int, map[string]interface{}) {
	var comments []models.Comment
	result := config.DB.Preload("Customer.User.Role").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&comments, "customer_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = comments
		return http.StatusOK, send
	}
}

func AddComment(comment models.Comment, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	var user models.User
	config.DB.Preload("Role").Find(&user, "id = ?", userId)
	if core.CUSTOMER != user.Role.Name {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	result := config.DB.Create(&comment)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Created)
		send["data"] = comment
		return http.StatusCreated, send
	}
}

func UpdateComment(comment models.Comment, key string, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)

	var c models.Comment
	config.DB.Find(&c, "id = ?", key)

	if userId != strconv.Itoa(c.StationeryId) {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	comment.Id, _ = strconv.Atoi(key)
	result := config.DB.Model(&models.Comment{}).Where("id = ?", key).Save(&comment)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Updated)
		send["data"] = comment
		return http.StatusCreated, send
	}
}
