package database

import (
	"fmt"

	"gorm.io/gorm"
)

type MysqlDB struct {
	*gorm.DB
}

func (m MysqlDB) dummy(a string){
	fmt.Println("dummy function")
}


