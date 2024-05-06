package database

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mseptian/qbit/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func ConnectDB() (*gorm.DB, error) {
	var err error
	godotenv.Load(".env")
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		return nil, errors.New("cant Connect to Database")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
	)

	return DB, nil
}
