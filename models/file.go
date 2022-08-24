package models

import "time"

type File struct {
	Id          int `gorm:"primaryKey"`
	CustomerId  int `gorm:"foreignKey"`
	Private     bool
	FilePath    string
	FolderId    string
	CreatedDate time.Time
	Customer    Customer
}
