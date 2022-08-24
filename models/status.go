package models

type Status struct {
	Id      int `gorm:"primaryKey"`
	Content string
}
