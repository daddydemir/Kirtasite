package models

type District struct {
	Id     int `gorm:"primaryKey"`
	Name   string
	CityId int `gorm:"foreignKey"`
}
