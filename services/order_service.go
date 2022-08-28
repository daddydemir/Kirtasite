package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"net/http"
)

func GetOrderById(key string) (int, map[string]interface{}) {
	var order models.Order
	result := config.DB.Preload("Customer.User.Role").Preload("File.Customer.Role").Preload("Status").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&order, "id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = order
		return http.StatusOK, send
	}
}

func GetOrdersByCustomerId(key string) (int, map[string]interface{}) {
	var orders []models.Order
	result := config.DB.Preload("Customer.User.Role").Preload("File.Customer.Role").Preload("Status").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&orders, "customer_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = orders
		return http.StatusOK, send
	}
}

func GetOrdersByStationeryId(key string) (int, map[string]interface{}) {
	var orders []models.Order
	result := config.DB.Preload("Customer.User.Role").Preload("File.Customer.Role").Preload("Status").Preload("Stationery.User.Role").Preload("Stationery.Address.City").Preload("Stationery.Address.District").Find(&orders, "stationery_id = ?", key)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = orders
		return http.StatusOK, send
	}
}

func AddOrder(order models.Order, token string) (int, map[string]interface{}) {
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
	result := config.DB.Create(&order)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Created)
		send["data"] = order
		return http.StatusCreated, send
	}
}

func UpdateOrder(order models.Order, key string, token string) (int, map[string]interface{}) {
	// customer iptal eder
	// stationery status güncellmesi yapar
	// bu endpoint kullanılmayacak
	result := config.DB.Model(&models.Order{}).Where("id = ?", key).Save(&order)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Updated)
		send["data"] = order
		return http.StatusOK, send
	}
}

func CancelOrder(key string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	var order models.Order
	config.DB.Find(&order, "id = ?", key)
	userId := auth.TokenParser(token)
	if userId != string(order.CustomerId) || order.StatusId != 1 {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}

	order.StatusId = 2 // iptal etme için gerekli olan code
	result := config.DB.Model(&models.Order{}).Where("id = ?", key).Save(&order)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Cancel)
		send["data"] = order
		return http.StatusOK, send
	}
}

func ConfirmOrder(key string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	var order models.Order
	config.DB.Find(&order, "id = ?", key)
	userId := auth.TokenParser(token)
	if userId != string(order.StationeryId) {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}

	order.StatusId = 3 // iptal etme için gerekli olan code
	result := config.DB.Model(&models.Order{}).Where("id = ?", key).Save(&order)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Confirm)
		send["data"] = order
		return http.StatusOK, send
	}
}

func ReadyOrder(key string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	var order models.Order
	config.DB.Find(&order, "id = ?", key)
	userId := auth.TokenParser(token)
	if userId != string(order.StationeryId) {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}

	order.StatusId = 4 // iptal etme için gerekli olan code
	result := config.DB.Model(&models.Order{}).Where("id = ?", key).Save(&order)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Ready)
		send["data"] = order
		return http.StatusOK, send
	}
}

func CompleteOrder(key string, token string) (int, interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	var order models.Order
	config.DB.Find(&order, "id = ?", key)
	userId := auth.TokenParser(token)
	if userId != string(order.StationeryId) {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}

	order.StatusId = 5 // iptal etme için gerekli olan code
	result := config.DB.Model(&models.Order{}).Where("id = ?", key).Save(&order)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Complate)
		send["data"] = order
		return http.StatusOK, send
	}
}
