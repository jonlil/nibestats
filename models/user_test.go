package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_ComparePassword(t *testing.T) {
	user := &User{
		Email: "tester@testson.com",
	}

	password := []byte("testpassword")

	user.SetPassword(password)
	assert.Equal(t, true, comparePasswords(user.Password, password))
	assert.Equal(t, true, user.Authenticate(password))
}
