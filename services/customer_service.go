package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"Kirtasite/secuirty"
	"fmt"
	"net/http"
)

func GetCustomers() (int, map[string]interface{}) {
	var customers []models.Customer
	result := config.DB.Preload("User.Role").Find(&customers)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = customers
		return http.StatusOK, send
	}
}

func GetCustomerByUserId(key string) (int, map[string]interface{}) {
	var customer models.Customer
	result := config.DB.Preload("User.Role").Find(&customer, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = customer
		return http.StatusOK, send
	}
}

func AddCustomer(customers models.Customer) (int, map[string]interface{}) {
	// TODO auth
	customers.User.Password = secuirty.HashPassword(customers.User.Password)
	fmt.Println("PAROLA : ", customers.User.Password)
	if customers.Username == "" {
		return http.StatusBadRequest, SendMessage(BadRequest)
	}
	result := config.DB.Create(&customers)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Created)
		send["data"] = customers
		return http.StatusCreated, send
	}
}

func UpdateCustomer(customers models.Customer, key string) (int, map[string]interface{}) {
	// TODO auth
	result := config.DB.Save(&customers)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		send := SendMessage(Updated)
		send["data"] = customers
		return http.StatusCreated, send
	}
}
