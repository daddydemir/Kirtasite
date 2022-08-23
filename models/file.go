package models

import "time"

type Files struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	UserId      int       `json:"user_id" gorm:"foreignKey"`
	Private     bool      `json:"private"`
	FilePath    string    `json:"file_path"`
	FolderId    string    `json:"folder_id"`
	CreatedDate time.Time `json:"created_date"`
	Customer    Customers `json:"customer" `
}
