package services

import (
	"errors"
	"ghgist-blog/models"
	"ghgist-blog/repository"
)

type FetchUserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewFetchUserService(userRepo repository.UserRepository) FetchWritersInterface {
	return &FetchUserServiceImpl{userRepo: userRepo}
}
func (s *FetchUserServiceImpl) FetchAllWriters() ([]models.User, error) {
	//fetch all users
	users, err := s.userRepo.FetchAllWriters()
	if err != nil {
		return nil, errors.New("could not fetch users")
	}
	return users, nil

}
