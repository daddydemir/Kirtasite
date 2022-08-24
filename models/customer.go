package models

type Customer struct {
	UserId   int `gorm:"primaryKey"`
	Username string
	User     User
}
