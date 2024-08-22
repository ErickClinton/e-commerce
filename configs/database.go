package configs

import (
	"eccomerce/internal/v1/user/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "host=localhost user=admin password=123 dbname=eccomercedb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}
