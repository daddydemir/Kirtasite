package models

import "time"

type Orders struct {
	Id             int           `json:"id"`
	FileId         int           `json:"file_id"`
	CustomerId     int           `json:"customer_id"`
	StationeryId   int           `json:"stationery_id"`
	StatusId       int           `json:"status_id"`
	TotalPrice     float32       `json:"total_price"`
	CreatedDate    time.Time     `json:"created_date"`
	DeliveryDate   time.Time     `json:"delivery_date"`
	FileData       Files         `json:"file" gorm:"foreignKey:file_id;references:id"`
	CustomerData   Customers     `json:"customer" gorm:"foreignKey:customer_id;references:user_id"`
	StationeryData Stationeryies `json:"stationery" gorm:"foreignKey:stationery_id;references:user_id"`
}
