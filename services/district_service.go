package services

import "Kirtasite/models"

func GetDistricts() (int, map[string]interface{}) {
	var districts []models.District
	return all(districts)
}

func GetDistrictByCityId(key string) (int, map[string]interface{}) {
	var district []models.District
	return query("city_id = ?", key, district)
}

func GetDistrictById(key string) (int, map[string]interface{}) {
	var district models.District
	return query("id = ?", key, &district)
}
