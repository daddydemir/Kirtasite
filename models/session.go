package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Id     uuid.UUID `gorm:"primaryKey"`
	Email  string
	CrDate time.Time
	ExDate time.Time
	Token  string
	Ip     string
}
