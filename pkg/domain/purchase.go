package domain

import (
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	UserID    string
	CofeeType string
	Time      time.Time
}
