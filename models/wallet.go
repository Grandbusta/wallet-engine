package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Email          string  `gorm:"unique;not null" json:"email"`
	Wallet_address int     `gorm:"not null" json:"wallet_address"`
	Amount         float64 `gorm:"default:0.00"`
}
