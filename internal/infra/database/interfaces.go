package database

import "github.com/sidney-cardoso/ecommerce-GO/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
