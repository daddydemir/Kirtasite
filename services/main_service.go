package services

import (
	"Kirtasite/config"
	"Kirtasite/core"
	"Kirtasite/models"
	"net/http"
)

func all(obj interface{}) (int, map[string]interface{}) {
	result := config.DB.Find(&obj)
	if result.Error != nil {
		return http.StatusNoContent, core.SendMessage(core.NoContent)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = obj
		return http.StatusOK, send
	}
}

func query(query string, key string, obj interface{}) (int, map[string]interface{}) {
	result := config.DB.Find(&obj, query, key)
	if result.Error != nil {
		return http.StatusNoContent, core.SendMessage(core.NoContent)
	} else {
		send := core.SendMessage(core.Ok)
		send["data"] = obj
		return http.StatusOK, send
	}
}

func save(addresses models.Address) (int, map[string]interface{}) {
	result := config.DB.Save(&addresses)
	if result.Error != nil {
		return http.StatusBadRequest, core.SendMessage(core.BadRequest)
	} else {
		msg := core.SendMessage(core.Updated)
		msg["data"] = addresses
		return http.StatusCreated, msg
	}
}
