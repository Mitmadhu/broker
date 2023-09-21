package model

import "gorm.io/gorm"

type Friend struct{
	gorm.Model
	UserName string
}

func (f Friend) CreateTable(db *gorm.DB) error{
	db.Exec("DROP TABLE IF EXISTS friends")
	return db.AutoMigrate(&Friend{})
}