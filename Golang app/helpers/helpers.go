package helpers

import (
	"log"
	"strings"
	"github.com/onspeedhp/golang_app/models"
	"github.com/onspeedhp/golang_app/database"
)

func CheckUserPass(username, password string) bool {
	// get userpass from database
	db := database.DB
	var users models.User
	// get user where name = username and password = password
	if err := db.Where("username = ? AND password = ?", username, password).First(&users).Error; err != nil {
		log.Println("error", err)
		return false
	} else {
		log.Println("success", users)
		return true
	}
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}