package models

type Districts struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	CityId int    `json:"city_id"`
}
