package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model     `json:"-"`
	Email          string  `gorm:"unique;not null" json:"email"`
	Wallet_address int     `gorm:"not null" json:"wallet_address"`
	Amount         float64 `gorm:"default:0.00" json:"amount"`
	Is_Active      bool    `gorm:"default:true" json:"is_active"`
}
