package models

import (
    "gorm.io/gorm"
)

type Repository struct {
    ID       		uint     `gorm:"primaryKey"`
    Name     		string
	Description 	string
    Commits  		[]Commit `gorm:"foreignKey:RepositoryID"`
    UserID   		uint
	Users       	[]User  `gorm:"many2many:repository_users;"`
}

func FindRepositoriesByUserId(db *gorm.DB, user *User) ([]Repository, error) {
	var repositories []Repository
	err := db.Where("user_id = ?",user.ID).First(&repositories).Error
	for i := range repositories {
		db.Model(&repositories[i]).Association("Commits").Find(&repositories[i].Commits)
	}
	return repositories, err
}

func FindRepositoryByName(db *gorm.DB, name string) (Repository, error) {
	var repository Repository
	err := db.Where("name = ?",name).First(&repository).Error
	db.Model(&repository).Association("Commits").Find(&repository.Commits)
	for i := range repository.Commits {
		var user User
		db.Where("id = ?",repository.Commits[i].UserID).First(&user)
		repository.Users = append(repository.Users, user)
	}

	// remove duplicates
	keys := make(map[uint]bool)
	list := []User{}
	for _, entry := range repository.Users {
		if _, value := keys[entry.ID]; !value {
			keys[entry.ID] = true
			list = append(list, entry)
		}
	}
	repository.Users = list
	
	return repository, err
}

