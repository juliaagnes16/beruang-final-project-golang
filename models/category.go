package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string
	Expenses []Expense
	Budget []Budget
}