package model

import "gorm.io/gorm"

type Friend struct{
	gorm.Model
	UserName string
}

func (f Friend) CreateTable(db *gorm.DB) error{
	return db.AutoMigrate(&Friend{})
}