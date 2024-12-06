package services

import (
	"crypto/sha512"
	"fmt"
	"testMedods2/config"
	"testMedods2/internals/repository"
	"testMedods2/x/interfacesx"
)

type UserService interface {
	CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*interfacesx.UserData, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) CreateUserAccount(userRequest *interfacesx.UserRegistrationRequest) (*interfacesx.UserData, error) {
	userRequest.Password = generatePasswordHash(userRequest.Password)

	userData, err := us.userRepo.CreateUserAccount(userRequest)
	if err != nil {
		return nil, err
	}

	return &interfacesx.UserData{
		ID:       userData.ID,
		Email:    userData.Email,
		UserName: userData.UserName,
	}, nil
}

func generatePasswordHash(password string) string {
	hash := sha512.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(config.Salt)))

}
