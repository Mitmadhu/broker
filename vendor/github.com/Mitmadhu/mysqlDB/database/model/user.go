package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Name     string
	Age      uint16
}

func (u User) ValidateUser(db *gorm.DB, username string) (bool, error){
	// validate the user
	return true, nil
}

// GetUserByUsername return user details by username
func (u User) GetUserByUsername(db *gorm.DB, username string)(User, error) {
	// fetch user by username
	return User{
		Name: "madhubala",
		Age:  32,
	}, nil
}

// GetUserByID returns details of the user by id
func (u User) GetUserByID(db *gorm.DB, ID string) (User, error) {
	// fetch user by id

	return User{
		Name: "madhubala",
		Age:  32,
	}, nil
}

func (u User) CreateTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
