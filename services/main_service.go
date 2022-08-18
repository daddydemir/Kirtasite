package services

import (
	"Kirtasite/config"
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