package models

import "time"

type Comments struct {
	Id             int          `json:"id"`
	CustomerId     int          `json:"customer_id" gorm:"foreignKey"`
	StationeryId   int          `json:"stationery_id" gorm:"foreignKey"`
	Content        string       `json:"content"`
	CreatedDate    time.Time    `json:"created_date"`
	Score          float32      `json:"score"`
	Customers      Customers    `json:"customer" gorm:"foreignKey:customer_id;references:user_id"`
	StationeryData Stationeries `json:"stationery" gorm:"foreignKey:stationery_id;references:user_id"`
}
