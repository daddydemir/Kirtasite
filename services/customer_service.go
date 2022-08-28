package services

import (
	"Kirtasite/auth"
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"Kirtasite/secuirty"
	"fmt"
	"net/http"
	"strconv"
)

func GetCustomers() (int, map[string]interface{}) {
	var customers []models.Customer
	result := config.DB.Preload("User.Role").Find(&customers)
	if result.Error != nil {
		return http.StatusNoContent, core.SendMessage(core.NoContent)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = customers
		return http.StatusOK, send
	}
}

func GetCustomerByUserId(key string) (int, map[string]interface{}) {
	var customer models.Customer
	result := config.DB.Preload("User.Role").Find(&customer, "user_id = ?", key)
	if result.Error != nil {
		return http.StatusNoContent, core.SendMessage(core.NoContent)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = customer
		return http.StatusOK, send
	}
}

func AddCustomer(customers models.Customer) (int, map[string]interface{}) {
	// TODO auth
	customers.User.Password = secuirty.HashPassword(customers.User.Password)
	fmt.Println("PAROLA : ", customers.User.Password)
	if customers.Username == "" {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	}
	result := config.DB.Create(&customers)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Created)
		send["data"] = customers
		return http.StatusCreated, send
	}
}

func UpdateCustomer(customers models.Customer, key string, token string) (int, map[string]interface{}) {
	s, m := auth.IsValid(token)
	if !s {
		return http.StatusUnauthorized, m
	}
	userId := auth.TokenParser(token)
	if userId != key {
		return http.StatusForbidden, core.SendMessage(core.Forbidden)
	}
	var tempUser models.User
	config.DB.Find(&tempUser, "id = ?", key)
	customers.User = tempUser
	customers.UserId, _ = strconv.Atoi(key)
	result := config.DB.Model(&models.Customer{}).Where("user_id = ?", key).Save(&customers)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		send := core.SendMessage(core.Updated)
		send["data"] = customers
		return http.StatusCreated, send
	}
}
