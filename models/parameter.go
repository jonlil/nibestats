package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jonlil/nibe-go"
	"time"
)

// AccessToken -
type Parameter struct {
	gorm.Model
	nibe.Parameter
	User     User
	UserID   int64
	Measured time.Time
}
