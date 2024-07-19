package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `gorm:"size:255;not null;"`
	Description string    `gorm:"type:text;"`
	Price       float64   `gorm:"type:numeric;"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New()
	return
}
