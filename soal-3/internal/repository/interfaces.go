package repository

import (
	"context"
	"github.com/mseptian/qbit/internal/models"
)

type UserRepository interface {
	Save(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id string) (user *models.User, RowsAffected int64)
	GetByEmail(ctx context.Context, email string) (user *models.User, RowsAffected int64)
}

type ProductRepository interface {
	GetAll(ctx context.Context) (user []*models.Product, err error)
	GetBySearch(ctx context.Context, search string) (user []*models.Product, err error)
	GetById(ctx context.Context, id string) (user *models.Product, RowsAffected int64)
	UpdateMultiple(ctx context.Context, product *[]models.Product) error
}

type CartRepository interface {
	Save(ctx context.Context, cart *models.Cart) error
	GetByUserIdAndProductId(ctx context.Context, userId, productId string) (cart *models.Cart, RowsAffected int64)
	Update(ctx context.Context, userId, productId string, cart *models.Cart) error
	GetAll(ctx context.Context, userId string) (cart []*models.CartAll, err error)
	RemoveByUserId(ctx context.Context, userId string) error
}
