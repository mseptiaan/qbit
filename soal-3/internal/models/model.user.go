package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID    string         `gorm:"primaryKey" json:"user_id"`
	Email     string         `gorm:"not null" json:"email"`
	Password  string         `gorm:"not null" json:"password"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	uuidHash, _ := uuid.NewV7()
	user.UserID = uuidHash.String()
	return nil
}
