package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	db, err := gorm.Open(sqlite.Open("myDatabase.db"))

	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR")
	}

	DB = db
	fmt.Println("DATABASE CONNECTED")
}
