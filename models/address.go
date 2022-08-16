package models

type Addresses struct {
	Id           int       `json:"id"`
	CityId       int       `json:"city_id"`
	DistrictId   int       `json:"district_id"`
	Header       string    `json:"header"`
	X            string    `json:"x"`
	Y            string    `json:"y"`
	CityData     Cities    `json:"city" gorm:"foreignKey:city_id;references:id"`
	DistrictData Districts `json:"district" gorm:"foreignKey:district_id;references:id"`
}
