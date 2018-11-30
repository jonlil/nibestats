package main

import (
	"github.com/jonlil/nibestats/database"
	"github.com/jonlil/nibestats/models"
	"github.com/thedevsaddam/renderer"
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
	db := database.Open()
	defer db.Close()

	db.AutoMigrate(&models.AccessToken{})
	db.AutoMigrate(&models.User{}).AddUniqueIndex("idx_user_email", "email")
	server.DB = db

	server.Listen()
}
