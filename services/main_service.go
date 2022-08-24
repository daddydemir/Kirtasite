package services

import (
	"Kirtasite/config"
	"Kirtasite/models"
	"net/http"
)

func all(obj interface{}) (int, map[string]interface{}) {
	result := config.DB.Find(&obj)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = obj
		return http.StatusOK, send
	}
}

func query(query string, key string, obj interface{}) (int, map[string]interface{}) {
	result := config.DB.Find(&obj, query, key)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = obj
		return http.StatusOK, send
	}
}

func save(addresses models.Address) (int, map[string]interface{}) {
	result := config.DB.Save(&addresses)
	if result.Error != nil {
		return http.StatusBadRequest, SendMessage(BadRequest)
	} else {
		msg := SendMessage(Updated)
		msg["data"] = addresses
		return http.StatusCreated, msg
	}
}
