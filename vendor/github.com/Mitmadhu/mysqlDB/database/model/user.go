package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Mitmadhu/broker/constants"
	"github.com/Mitmadhu/mysqlDB/config"
	"github.com/Mitmadhu/mysqlDB/helper"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	FirstName string `gorm:"type:varchar(50)"`
	LastName  string `gorm:"type:varchar(50)"`
	Age       uint16
	Password  string `gorm:"not null"`
	Salt      string `gorm:"type:varchar(255);not null"`
}

func (u User) ValidateUser(username string, password string) (bool, error) {
	// validate the user
	user, err := User{}.GetUserByUsername(username)
	if err != nil {
		return false, errors.New(constants.UserNotFound)
	}
	password += user.Salt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil, nil
}

// GetUserByUsername return user details by username
func (u User) GetUserByUsername(username string) (*User, error) {
	// fetch user by username
	var user User
	db := config.GetDB()
	result := db.Where("Username = ?", username).First(&user)
	if result == nil {
		return nil, errors.New(constants.InternalServerError)
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil, errors.New(constants.UserNotFound)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &User{
		Username: user.Username,
		Age:      user.Age,
		Password: user.Password,
		Salt:     user.Salt,
	}, nil
}

// GetUserByID returns details of the user by id
func (u User) GetUserByID(ID string) (User, error) {
	// fetch user by id

	return User{
		Username: "madhubala",
		Age:      32,
	}, nil
}

func (u User) CreateTable(db *gorm.DB) error {
	db.Exec("DROP TABLE IF EXISTS users")
	return db.AutoMigrate(&User{})
}

func (u User) Register(username, password, firstName, lastName string, age uint16) error {
	salt := helper.GetSalt()
	password += salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := User{
		Username:  username,
		Password:  string(hashedPassword),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Salt:      salt,
	}
	result := config.GetDB().Create(&user)
	return result.Error
}

func (u User) DeleteUser(username string) error {
	result := config.GetDB().Where("Username = ?", username).Delete(&User{})
	return result.Error
}
