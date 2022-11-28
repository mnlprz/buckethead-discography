package database

import (
	"log"

	"github.com/mnlprz/buckethead-discography/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnectionDB() *gorm.DB {
	dsn := "host=localhost user=root password=secret dbname=buckethead port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SetUpDB(db *gorm.DB) {
	db.AutoMigrate(&models.Album{})
}
