package models

type City struct {
	Id   int `gorm:"primaryKey"`
	Name string
}
