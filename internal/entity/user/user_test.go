package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jhon Doe", "Jhon@email.com", "Mypass")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jhon Doe", user.Name)
	assert.Equal(t, "Jhon@email.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Jhon Doe", "Jhon@email.com", "Mypass")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("Mypass"))
	assert.False(t, user.ValidatePassword("password"))
	assert.NotEqual(t, "Mypass", user.Password)
}
