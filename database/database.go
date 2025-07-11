package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable to hold the database connection

func Connect() {

	db, err := gorm.Open(sqlite.Open("myDatabase.db"), &gorm.Config{}) // Open a connection to the SQLite database

	if err != nil {
		fmt.Println("DATABASE CONNECTION ERROR")
	}

	DB = db // Assign the database connection to the global variable
	fmt.Println("DATABASE CONNECTED")
}
