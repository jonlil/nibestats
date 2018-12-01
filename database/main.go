package database

import (
	"github.com/jinzhu/gorm"
	// Dialect import, not used directly
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

func Open() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Println(err)
		panic("failed connecting to database.")
	}
	return db
}
