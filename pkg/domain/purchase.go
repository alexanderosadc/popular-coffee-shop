package domain

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	UserID    uint
	CofeeType string
	Time      time.Time
}
