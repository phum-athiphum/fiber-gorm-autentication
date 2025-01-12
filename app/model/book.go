package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          int64  `gorm:"unique;not null" json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Author      string `json:"author" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"max=255"`
	Price       uint   `json:"price" validate:"required,gt=0"`
}
