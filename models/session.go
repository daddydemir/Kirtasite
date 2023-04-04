package models

import "time"

type Session struct {
	Id     int `gorm:"primaryKey"`
	Email  string
	CrDate time.Time
	ExDate time.Time
	Token  string
	Ip     string
}
