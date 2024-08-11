package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model
	UserId     uint    `json:"user_id"`
	CategoryID uint    `json:"category_id"`
	Amount     float64 `json:"amount"`
}
