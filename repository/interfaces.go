package repository

import "ghgist-blog/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	SetupIndexes() error
	FetchAllUsers() ([]models.User, error)
}
type ArticleRepository interface{}
