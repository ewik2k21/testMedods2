package repository

import (
	"testMedods2/internals/model"
	"testMedods2/x/interfacesx"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*model.User, error)
	GetUserByEmail(userEmail *string) (*model.User, error)
	UpdateRefreshTokenDb(email, userIP, refreshToken *string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*model.User, error) {
	user := &model.User{
		Email:        userRequest.Email,
		UserName:     userRequest.UserName,
		PasswordHash: userRequest.Password,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(userEmail *string) (*model.User, error) {
	user := &model.User{}

	if err := r.db.Where("email = ?", userEmail).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdateRefreshTokenDb(email, userIP, refreshToken *string) error {
	if err := r.db.Model(&model.User{}).Where("email = ?", email).UpdateColumn("refresh_token", refreshToken).UpdateColumn("user_ip", userIP).Error; err != nil {
		return err
	}
	return nil
}
