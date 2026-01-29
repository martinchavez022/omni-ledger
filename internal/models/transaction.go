package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount			float64 `json:"amount"`
	Description		string  `json:"description"`
	Category		string	`json:"category"`
}