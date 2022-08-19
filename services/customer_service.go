package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"fmt"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetCustomers() (int, map[string]interface{}) {
	var customers []models.Customers
	result := config.DB.Preload(clause.Associations).Find(&customers)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = customers
		return http.StatusOK, send
	}
}

func GetCustomerByUserId(key string) (int, map[string]interface{}) {
	var customer models.Customers
	result := config.DB.Preload(clause.Associations).Find(&customer, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = customer
		return http.StatusOK, send
	}
}

func AddCustomer(customers models.Customers) (int, map[string]interface{}) {
	// TODO auth
	fmt.Println("PAROLA : ", customers.UserData.Password)
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

func UpdateCustomer(customers models.Customers, key string) (int, map[string]interface{}) {
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
