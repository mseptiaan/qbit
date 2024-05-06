package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ProductId string         `gorm:"primaryKey" json:"product_id"`
	Name      string         `gorm:"not null" json:"name"`
	Qty       int            `gorm:"not null" json:"qty"`
	Price     int            `gorm:"not null" json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
