package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Mitmadhu/broker/constants"
	"github.com/Mitmadhu/mysqlDB/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"primaryKey"`
	Name     string
	Age      uint16
	Password string
}

func (u User) ValidateUser(db *gorm.DB, username string) (bool, error){
	// validate the user
	return true, nil
}

// GetUserByUsername return user details by username
func (u User) GetUserByUsername(db *gorm.DB, username string)(User, error) {
	// fetch user by username
	var user User;
	result := db.Where( "Username = ?", username).First(&user)
	if(result == nil){
		return User{}, errors.New(constants.InternalServerError)
	}
	if (result.Error == gorm.ErrRecordNotFound){
		return nil, errors.New(constants.UserNotFound)
	}
	if(result.Error != nil){
		return nil, result.Error
	}

	return User{
		Username: user.Username,
		Age: user.Age,
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
