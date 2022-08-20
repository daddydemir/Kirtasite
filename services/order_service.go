package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetOrderById(key string) (int, map[string]interface{}) {
	var order models.Orders
	result := config.DB.Preload(clause.Associations).Find(&order, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = order
		return http.StatusOK, send
	}
}

func GetOrdersByCustomerId(key string) (int, map[string]interface{}) {
	var orders []models.Orders
	result := config.DB.Preload(clause.Associations).Find(&orders, "customer_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = orders
		return http.StatusOK, send
	}
}

func GetOrdersByStationeryId(key string) (int, map[string]interface{}) {
	var orders []models.Orders
	result := config.DB.Preload(clause.Associations).Find(&orders, "stationery_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = orders
		return http.StatusOK, send
	}
}

func AddOrder(order models.Orders) (int, map[string]interface{}) {
	result := config.DB.Create(&order)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = order
		return http.StatusCreated, send
	}
}

func UpdateOrder(order models.Orders, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&order)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = order
		return http.StatusOK, send
	}
}
