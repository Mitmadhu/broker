package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func DBFactory(driver, dns string) (*gorm.DB, error) {
	dsn := "madhu:mad.ayush@tcp(127.0.0.1:3306)/microservice?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
