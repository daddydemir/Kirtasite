package models

type Districts struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	CityId int    `json:"city_id" gorm:"foreignKey"`
}
