package auth

import (
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"Kirtasite/secuirty"
	"net/http"
	"strconv"
)

func Login(email string, password string) (int, interface{}) {
	var user models.User
	config.DB.Preload("Role").Find(&user, "mail = ?", email)
	if user.Id == 0 {
		return http.StatusBadRequest, core.SendMessage(core.LoginFail)
	}
	if user.Role.Name == core.STATIONERY {
		var stationery models.Stationery
		config.DB.Preload("Address.District").Preload("User.Role").Preload("Address.City").Find(&stationery, "user_id = ? ", user.Id)
		isTrue := secuirty.CheckPassword(stationery.User.Password, password)
		if isTrue {
			id := strconv.Itoa(user.Id)
			token := GenerateToken(id)
			send := core.SendMessage(core.Ok)
			send["data"] = stationery
			send["token"] = token
			return http.StatusOK, send
		} else {
			return http.StatusBadRequest, core.SendMessage(core.LoginFail)
		}
	} else if user.Role.Name == core.CUSTOMER {
		var customer models.Customer
		config.DB.Preload("User.Role").Find(&customer, "user_id = ?", user.Id)
		isTrue := secuirty.CheckPassword(customer.User.Password, password)
		if isTrue {
			id := strconv.Itoa(user.Id)
			token := GenerateToken(id)
			send := core.SendMessage(core.Ok)
			send["data"] = customer
			send["token"] = token
			return http.StatusOK, send
		} else {
			return http.StatusBadRequest, core.SendMessage(core.LoginFail)
		}
	}
	return http.StatusBadRequest, core.SendMessage(core.LoginFail)
}
