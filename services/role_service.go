package services

import (
	"Kirtasite/models"
)

func GetRoles() (int, map[string]interface{}) {
	var roles []models.Roles
	return all(roles)
	/*result := config.DB.Find(&roles)
	if result.Error != nil {
		return http.StatusNoContent, SendMessage(NoContent)
	} else {
		send := SendMessage(Ok)
		send["data"] = roles
		return http.StatusOK, send
	}*/
}
