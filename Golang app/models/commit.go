package models

// import (
//     "gorm.io/gorm"
// )

type Commit struct {
    ID           uint   `gorm:"primaryKey"`
    Message      string
    RepositoryID uint
    UserID       uint
}

// // Find commits by user id
// func FindCommitsByUserId(db *gorm.DB, User *User) ([]Commit, error) {
// 	var commits []Commit
// 	err := db.Where("userId = ?", User.ID).Find(&commits).Error
// 	return commits, err
// }

// // Find commits by repository id
// func FindCommitsByRepositoryId(db *gorm.DB, Repository *Repository) ([]Commit, error) {
// 	var commits []Commit
// 	err := db.Where("repositoryId = ?", Repository.ID).Find(&commits).Error
// 	return commits, err
// }




