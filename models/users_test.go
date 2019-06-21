package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_UserStore(t *testing.T) {
	UserStore = make(map[string]User)
	// Check if user store is empty
	assert.Empty(t, UserStore)

	// Add user to user store
	UserStore = map[string]User{
		"1": User{
			ID:        1,
			Firstname: "Abubakar",
			Lastname:  "Oladeji",
		},
	}

	// Verify user store count is greater than zero
	assert.True(t, len(UserStore) > 0)
}

func TestUsers_GeneratingUserID(t *testing.T) {
	UserStore = make(map[string]User)
	assert.Empty(t, UserStore)

	// Empty user should return 1
	assert.Equal(t, 1, GenerateUserID())

	// Add user to user store
	UserStore = map[string]User{
		"1": User{
			ID:        1,
			Firstname: "Abubakar",
			Lastname:  "Oladeji",
		},
	}

	assert.Equal(t, 2, GenerateUserID())
}

func TestUsers_OK(t *testing.T) {
	UserStore = make(map[string]User)
	assert.Empty(t, UserStore)
	user := User{}

	err := user.OK()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "first_name")

	user.Firstname = "Abubakar"
	err = user.OK()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "last_name")

	user.Lastname = "Oladeji"
	err = user.OK()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "email")

	user.Email = "abc@abc.com"
	err = user.OK()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "password")

	user.Password = "password"
	err = user.OK()
	assert.Nil(t, err)
}
