package models

import "time"

type Orders struct {
	Id           int          `json:"id" gorm:"primaryKey"`
	FileId       int          `json:"file_id" gorm:"foreignKey"`
	CustomerId   int          `json:"customer_id" gorm:"foreignKey"`
	StationeryId int          `json:"stationery_id" gorm:"foreignKey"`
	StatusId     int          `json:"status_id" gorm:"foreignKey"`
	TotalPrice   float32      `json:"total_price"`
	CreatedDate  time.Time    `json:"created_date"`
	DeliveryDate time.Time    `json:"delivery_date"`
	File         Files        `json:"file"`
	Customer     Customers    `json:"customer"`
	Stationery   Stationeries `json:"stationery"`
	Status       Status       `json:"status"`
}
