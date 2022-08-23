package models

type Prices struct {
	Id           int     `json:"id" gorm:"primaryKey"`
	StationeryId int     `json:"stationery_id" gorm:"foreignKey"`
	Info         string  `json:"info"`
	Price        float32 `json:"price"`
}
