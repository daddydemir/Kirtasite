package models

type User struct {
	Id        int `gorm:"primaryKey"`
	RoleId    int `gorm:"foreignKey"`
	Password  string
	ImagePath string
	Phone     string
	Mail      string
	Role      Role
}
