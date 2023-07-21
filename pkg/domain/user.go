package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              string
	Membership      string
	PurchaseHistory []Purchase `gorm:"foreignKey:UserID"`
}
