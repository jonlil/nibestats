package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// User -
type User struct {
	gorm.Model
	Email    string
	Name     string
	Password string
}

// SetPassword -
func (u *User) SetPassword(rawPassword []byte) {
	hash := hashAndSalt(rawPassword)

	u.Password = hash
}

// Authenticate -
func (u *User) Authenticate(plainPwd []byte) bool {
	return comparePasswords(u.Password, plainPwd)
}

func hashAndSalt(pwd []byte) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
