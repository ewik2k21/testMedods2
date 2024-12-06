package repository

import (
	"testMedods2/internals/model"
	"testMedods2/x/interfacesx"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*model.User, error) {

}
