package repository

import (
	"context"
	"github.com/mseptian/qbit/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, user *models.User) error {
	result := r.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (user *models.User, RowsAffected int64) {
	tx := r.db.WithContext(ctx).First(&user, "user_id = ?", id)
	return user, tx.RowsAffected
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (user *models.User, RowsAffected int64) {
	tx := r.db.WithContext(ctx).First(&user, "email = ?", email)
	return user, tx.RowsAffected
}

func (r *userRepository) GetAll(ctx context.Context) (user []*models.User, err error) {
	return user, err
}

func (r *userRepository) Update(ctx context.Context, user bson.M) error {
	return nil
}

func (r *userRepository) DeleteById(ctx context.Context, id string) error {
	return nil
}
