package repository

import (
	"context"
	"fmt"
	"github.com/mseptian/qbit/internal/models"
	"gorm.io/gorm"
	"strings"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(ctx context.Context) (product []*models.Product, err error) {
	tx := r.db.WithContext(ctx).Model(&models.Product{}).Find(&product)
	return product, tx.Error
}

func (r *productRepository) GetBySearch(ctx context.Context, search string) (product []*models.Product, err error) {
	tx := r.db.WithContext(ctx).Model(&models.Product{}).Where("LOWER(name) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(search))).Find(&product)
	return product, tx.Error
}

func (r *productRepository) GetById(ctx context.Context, id string) (product *models.Product, RowsAffected int64) {
	tx := r.db.WithContext(ctx).Model(&models.Product{}).Where("product_id = ?", id).First(&product)
	return product, tx.RowsAffected
}

func (r *productRepository) UpdateMultiple(ctx context.Context, product *[]models.Product) error {
	tx := r.db.WithContext(ctx).Save(product)
	return tx.Error
}
