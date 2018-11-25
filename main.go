package main

import (
	"github.com/jinzhu/gorm"
	// Dialect import, not used directly
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jonlil/nibestats/models"
	"github.com/thedevsaddam/renderer"
	"log"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}

	rnd = renderer.New(opts)
}

func main() {
	server := NewServer()

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Println(err)
		panic("failed connecting to database.")
	}
	defer db.Close()

	db.AutoMigrate(&models.AccessToken{})
	db.AutoMigrate(&models.User{}).AddUniqueIndex("idx_user_email", "email")
	server.DB = db

	server.Listen()
}
