package main

import (
  "github.com/jinzhu/gorm"
  // Dialect import, not used directly
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "log"
  "./nibestats"
)

func main() {
  server := nibestats.NewServer()

  db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    log.Println(err)
    panic("failed connecting to database.")
  }
  defer db.Close()

  db.AutoMigrate(&nibestats.AccessToken{})

  server.DB = db

  server.Listen()
}
