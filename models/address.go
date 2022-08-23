package models

type Addresses struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	CityId     int       `json:"city_id" gorm:"foreignKey"`
	DistrictId int       `json:"district_id" gorm:"foreignKey"`
	Header     string    `json:"header"`
	X          string    `json:"x"`
	Y          string    `json:"y"`
	City       Cities    `json:"city"`
	District   Districts `json:"district"`
}
