package models

type Address struct {
	Id         int `gorm:"primaryKey"`
	CityId     int `gorm:"foreignKey"`
	DistrictId int `gorm:"foreignKey"`
	Header     string
	X          string
	Y          string
	City       City
	District   District
}
