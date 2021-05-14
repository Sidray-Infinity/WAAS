package Model

import (
	entity "waas/Model/entity"
)

type UserModel interface {
	GetUser(userId int) (*entity.User, error)
	RegisterUser(newUser entity.User) error
	ValidateUsername(newUser entity.User) bool
	ValidateKYC(newUser entity.User) bool
}
