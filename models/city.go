package models

type Cities struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
