package models

import (
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	UserId    string `gorm:"not null" json:"user_id"`
	ProductId string `gorm:"not null" json:"product_id"`
	Qty       int    `gorm:"not null" json:"qty"`
}

type CartAll struct {
	ProductId  string `json:"product_id"`
	Name       string `json:"name"`
	CartQty    int    `json:"cart_qty"`
	ProductQty int    `json:"product_qty"`
	Price      int    `json:"price"`
}

func (cart *Cart) BeforeCreate(tx *gorm.DB) error {
	var product Product
	trx := tx.Model(&Product{}).Where("product_id = ?", cart.ProductId).First(&product)
	if trx.RowsAffected < 1 {
		return errors.New("product not found")
	}
	if (product.Qty - cart.Qty) < 0 {
		return errors.New("insufficient qty")
	}
	return nil
}

func (cart *Cart) BeforeUpdate(tx *gorm.DB) (err error) {
	var product Product
	trx := tx.Model(&Product{}).Where("product_id = ?", cart.ProductId).First(&product)
	if trx.RowsAffected < 1 {
		return errors.New("product not found")
	}
	if (product.Qty - cart.Qty) < 0 {
		return errors.New("insufficient qty")
	}
	return nil
}
