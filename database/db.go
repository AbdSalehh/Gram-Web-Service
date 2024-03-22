package database

import (
	"MyGram/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = "localhost"
	user = "postgres"
	password = "saleh387"
	dbPort   = "5433"
	dbname   = "postgres"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to database")
	db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}