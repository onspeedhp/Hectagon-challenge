package database

import (
	"github.com/onspeedhp/golang_app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:0931588846hP@/git_go_auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	DB.AutoMigrate(&models.User{}, &models.Repository{}, &models.Commit{})
}