package main

import (
	"github.com/jinzhu/gorm"
	"github.com/jonlil/nibe-go"
	"github.com/jonlil/nibestats/database"
	"github.com/jonlil/nibestats/models"
	"log"
	"time"
)

var (
	nibeAPI *nibe.API
	db      *gorm.DB
)

func getTwentyMinutesAgo() time.Time {
	return time.Now().Add(time.Duration(-5) * time.Minute)
}

func getRefreshTokens() []models.AccessToken {
	var tokens []models.AccessToken
	db.Where("updated_at < ?", getTwentyMinutesAgo()).Find(&tokens)

	return tokens
}

func main() {
	db = database.Open()
	for {
		log.Println("Refreshing tokens older then 20 minutes")
		time.Sleep(time.Duration(5) * time.Second)

		for _, token := range getRefreshTokens() {
			err := token.Refresh()
			if err != nil {
				log.Println("Failed refreshing token for user")
				log.Println(err)
				continue
			}
			db.Save(&token)
		}
	}
}
