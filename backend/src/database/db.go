package database

import (
	"ambassador/src/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(mysql.Open("user:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the database!")
	}
	fmt.Println("Database connection successfully opened!")
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
