package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jonlil/nibe-go"
)

// AccessToken -
type AccessToken struct {
	gorm.Model
	nibe.AccessToken
	User   User
	UserID int64
}
