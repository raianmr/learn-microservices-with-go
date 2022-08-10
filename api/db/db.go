package db

import (
	"fmt"
	"log"

	"github.com/raianmr/learn-microservices-with-go/api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DefaultDB() *gorm.DB {
	e := config.DefaultEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", e.DB_HOST, e.DB_USER, e.DB_PASS, e.DB_NAME, e.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&User{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
