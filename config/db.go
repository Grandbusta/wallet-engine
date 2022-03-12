package config

import (
	"fmt"
	"log"
	"os"
	"wallet-engine/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(mysql.Open(DATABASE_URL), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to DB")
	} else {
		fmt.Println("connected to DB")
	}
	database.Debug().AutoMigrate(
		&models.Wallet{},
	)

	DB = database
}
