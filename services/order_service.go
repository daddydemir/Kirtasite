package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"net/http"
)

func GetOrderById(key string) (int, map[string]interface{}) {
	var order models.Order
	result := config.DB.Preload("Customer.User.Role").Preload("File.Customer.Role").Preload("Status").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&order, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = order
		return http.StatusOK, send
	}
}

func GetOrdersByCustomerId(key string) (int, map[string]interface{}) {
	var orders []models.Order
	result := config.DB.Preload("Customer.User.Role").Preload("File.Customer.Role").Preload("Status").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&orders, "customer_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = orders
		return http.StatusOK, send
	}
}

func GetOrdersByStationeryId(key string) (int, map[string]interface{}) {
	var orders []models.Order
	result := config.DB.Preload("Customer.User.Role").Preload("File.Customer.Role").Preload("Status").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&orders, "stationery_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Ok)
		send["data"] = orders
		return http.StatusOK, send
	}
}

func AddOrder(order models.Order) (int, map[string]interface{}) {
	result := config.DB.Create(&order)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = order
		return http.StatusCreated, send
	}
}

func UpdateOrder(order models.Order, key string) (int, map[string]interface{}) {
	result := config.DB.Save(&order)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = order
		return http.StatusOK, send
	}
}
