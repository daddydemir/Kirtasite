package models

type Roles struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
