package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_UserStore(t *testing.T) {
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
