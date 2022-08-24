package services

import "Kirtasite/models"

func GetStatuses() (int, map[string]interface{}) {
	var statuses []models.Status
	return all(statuses)
}
func GetStatusById(key string) (int, map[string]interface{}) {
	var status models.Status
	return query("id = ?", key, &status)
}
