package database

import (
	"fmt"
	"go-iris-template/config"
	"go-iris-template/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	connectionStr := config.GetDatabaseConfig()
	fmt.Println(connectionStr)
	DB, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}
}
