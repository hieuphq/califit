package repository

import "github.com/hieuphq/califit/src/domain/model"

// UserRepository user repository
type UserRepository interface {
	FindAll(repo DBRepo, param QueryParam) ([]model.User, error)
	FindByEmail(repo DBRepo, email string) (*model.User, error)
	Save(repo DBRepo, user *model.User) error
}
