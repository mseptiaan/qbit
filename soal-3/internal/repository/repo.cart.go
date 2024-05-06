package repository

import (
	"context"
	"github.com/mseptian/qbit/internal/models"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) Save(ctx context.Context, cart *models.Cart) error {
	tx := r.db.WithContext(ctx).Create(cart)
	return tx.Error
}

func (r *cartRepository) GetByUserIdAndProductId(ctx context.Context, userId, productId string) (cart *models.Cart, RowsAffected int64) {
	tx := r.db.WithContext(ctx).First(&cart, "user_id = ? AND product_id = ?", userId, productId)
	return cart, tx.RowsAffected
}

func (r *cartRepository) Update(ctx context.Context, userId, productId string, cart *models.Cart) error {
	tx := r.db.WithContext(ctx).Model(&cart).Where("user_id = ? AND product_id = ?", userId, productId).Updates(&cart)
	return tx.Error
}

func (r *cartRepository) GetAll(ctx context.Context, userId string) (cart []*models.CartAll, err error) {
	tx := r.db.WithContext(ctx).Raw("SELECT "+
		"p.product_id product_id, p.name name, c.qty cart_qty, p.qty product_qty, p.price price "+
		"FROM carts c "+
		"INNER JOIN products p "+
		"ON c.product_id=p.product_id "+
		"WHERE c.user_id = ?", userId).Scan(&cart)
	return cart, tx.Error
}

func (r *cartRepository) RemoveByUserId(ctx context.Context, userId string) error {
	tx := r.db.WithContext(ctx).Unscoped().Where("user_id = ?", userId).Delete(&models.Cart{})
	return tx.Error
}
