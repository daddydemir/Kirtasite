package services

import "Kirtasite/models"

func GetOrderStatuses() (int, map[string]interface{}) {
	var statuses []models.Order_Statuses
	return all(statuses)
}
func GetOrderStatusById(key string) (int, map[string]interface{}) {
	var status models.Order_Statuses
	return query("id = ?", key, &status)
}
