package models

import "time"

type Comment struct {
	Id           int `gorm:"primaryKey"`
	CustomerId   int `gorm:"foreignKey"`
	StationeryId int `gorm:"foreignKey"`
	Content      string
	CreatedDate  time.Time
	Score        float32
	Customer     Customer
	Stationery   Stationery
}
