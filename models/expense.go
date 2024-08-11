package models

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	CategoryID uint    `json:"category_id"`
	Amount     float64 `json:"amount"`
	Note       string  `json:"note"`
}
