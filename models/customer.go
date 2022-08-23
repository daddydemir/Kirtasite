package models

type Customers struct {
	UserId   int    `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username"`
	User     Users  `json:"user"`
}
