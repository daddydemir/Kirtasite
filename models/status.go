package models

type Status struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
}
