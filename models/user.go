package models

type Users struct {
	Id        int    `json:"id"`
	RoleId    int    `json:"role_id"`
	Password  string `json:"-"`
	ImagePath string `json:"image_path"`
	Phone     string `json:"phone"`
	RoleDate  Roles  `json:"role" gorm:"foreignKey:role_id;references:id"`
}
