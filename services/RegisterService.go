package services

import (
	"AuthGo/models"
	"AuthGo/repository"
	"AuthGo/utils"
	"AuthGo/validation"
)

// this stores all functions from the user repository
type RegisterServiceImpl struct {
	userRepo repository.UserRepository
}

func NewRegisterService(userRepo repository.UserRepository) RegisterInterface {
	return &RegisterServiceImpl{userRepo: userRepo}
}

func (s *RegisterServiceImpl) RegisterUser(email, username, password, role string) (*models.User, error) {
	//validating the input calling the validation functions

	err := validation.ValidateEmail(email)
	if err != nil {
		return nil, err
	}
	err = validation.ValidatePassword(password)
	if err != nil {
		return nil, err
	}
	err = validation.ValidateUsername(username)
	if err != nil {
		return nil, err
	}

	//hashing password calling the utils function
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	//creating user body calling the models struct

	user := &models.User{
		Email:    email,
		Password: hashedPassword,
		Username: username,
		Role:     role,
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
