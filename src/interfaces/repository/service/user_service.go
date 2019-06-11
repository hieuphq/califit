package service

import (
	"fmt"

	"github.com/hieuphq/califit/src/interfaces/repository"
	"github.com/jinzhu/gorm"
)

// UserService service to work with User
type UserService struct {
	repo repository.UserRepository
}

// NewUserService make a new user service instance
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Duplicated validate a email is existed in persistence data
func (s *UserService) Duplicated(repo repository.DBRepo, email string) error {
	user, err := s.repo.FindByEmail(repo, email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if user.ID > 0 {
		return fmt.Errorf("%s already exists", email)
	}

	return nil
}
