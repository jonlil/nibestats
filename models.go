package nibestats

import (
	"github.com/jinzhu/gorm"
)

// AccessToken -
type AccessToken struct {
	gorm.Model
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}