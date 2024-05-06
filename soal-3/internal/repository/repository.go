package repository

import (
	"gorm.io/gorm"
)

type Repositories struct {
	User    UserRepository
	Product ProductRepository
	Cart    CartRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User:    NewUserRepository(db),
		Product: NewProductRepository(db),
		Cart:    NewCartRepository(db),
	}
}
