package database

import (
	"testing"

	"github.com/sidney-cardoso/ecommerce-GO/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Sidney", "sidney@email.com", "123456")
	userDB := NewUser(db)

	err = userDB.CreateUser(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound).Where("id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, user.Password)
}
