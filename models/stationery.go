package models

type Stationery struct {
	UserId      int `gorm:"primaryKey"`
	AddressId   int `gorm:"foreignKey"`
	CompanyName string
	Score       float32
	User        User
	Address     Address
	Prices      []Price
}
