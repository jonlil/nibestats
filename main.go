package main

import (
  "./nibestats"
  "github.com/jinzhu/gorm"
)

func main() {
  server := nibestats.NewServer()

  db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("failed connecting to database.")
  }
  server.DB = db
  defer db.Close()
}
