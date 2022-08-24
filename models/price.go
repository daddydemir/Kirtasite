package models

type Price struct {
	Id           int `gorm:"primaryKey"`
	StationeryId int `gorm:"foreignKey"`
	Info         string
	Price        float32
}
