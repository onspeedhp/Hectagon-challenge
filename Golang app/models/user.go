package models

import (
    "gorm.io/gorm"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Name     string
    Password string
    Repos    []Repository `gorm:"foreignKey:UserID"`
    Commits  []Commit     `gorm:"foreignKey:UserID"`
}

func FindUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	err := db.Where("username = ?",username).First(&user).Error
	db.Model(&user).Association("Repos").Find(&user.Repos)
	db.Model(&user).Association("Commits").Find(&user.Commits)
	return user, err
}