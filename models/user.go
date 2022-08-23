package models

type Users struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	RoleId    int    `json:"role_id" gorm:"foreignKey"`
	Password  string `json:"-"`
	ImagePath string `json:"image_path"`
	Phone     string `json:"phone"`
	Mail      string `json:"mail"`
	Role      Roles  `json:"role"`
}
