package models

import "time"

type Order struct {
	Id           int `gorm:"primaryKey"`
	FileId       int `gorm:"foreignKey"`
	CustomerId   int `gorm:"foreignKey"`
	StationeryId int `gorm:"foreignKey"`
	StatusId     int `gorm:"foreignKey"`
	TotalPrice   float32
	CreatedDate  time.Time
	DeliveryDate time.Time
	File         File
	Customer     Customer
	Stationery   Stationery
	Status       Status
}
