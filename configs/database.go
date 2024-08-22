package configs

import (
	"eccomerce/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "host=localhost user=admin password=123 dbname=eccomercedb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	for _, model := range utils.AllModels() {
		if err := db.AutoMigrate(model); err != nil {
			panic("failed to migrate database: " + err.Error())
		}
	}
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}
