package services

import (
	"Kirtasite/models"
)

func GetRoles() (int, map[string]interface{}) {
	var roles []models.Role
	return all(roles)
}
