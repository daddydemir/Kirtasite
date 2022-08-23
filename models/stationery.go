package models

type Stationeries struct {
	UserId      int       `json:"user_id" gorm:"primaryKey"`
	CompanyName string    `json:"company_name"`
	AddressId   int       `json:"address_id" gorm:"foreignKey"`
	Score       float32   `json:"score"`
	User        Users     `json:"user"`
	Address     Addresses `json:"address"`
	Prices      []*Prices
}
