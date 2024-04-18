package configs

import (
	"fmt"
	"log"

	"github.com/altsaqif/go-restapi-mux/cmd/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DB_USER := GoDotEnvVariable("DB_USER")
	DB_PASSWORD := GoDotEnvVariable("DB_PASSWORD")
	DB_HOST := GoDotEnvVariable("DB_HOST")
	DB_PORT := GoDotEnvVariable("DB_PORT")
	DB_NAME := GoDotEnvVariable("DB_NAME")

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database Connection Fail! %v", err.Error())
	}
	err = db.AutoMigrate(&models.Product{}, &models.Users{})
	if err != nil {
		log.Println("Failed to auto migrate")
	}
	DB = db
}
