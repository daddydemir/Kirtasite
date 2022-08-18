package services

import "Kirtasite/models"

func GetDistricts() (int, map[string]interface{}) {
	var districts []models.Districts
	return all(districts)
}

func GetDistrictByCityId(key string) (int, map[string]interface{}) {
	var district []models.Districts
	return query("city_id = ?", key, district)
}

func GetDistrictById(key string) (int, map[string]interface{}) {
	var district models.Districts
	return query("id = ?", key, &district)
}
