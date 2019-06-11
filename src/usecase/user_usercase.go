package usecase

import (
	"github.com/hieuphq/califit/src/domain/model"
	"github.com/hieuphq/califit/src/interfaces/repository"
	"github.com/hieuphq/califit/src/interfaces/repository/service"
)

// UserUsecase usecase for user entity
type UserUsecase interface {
	ListUser() ([]model.User, error)
	RegisterUser(input model.User) (*model.User, error)
}

type userUsecase struct {
	repo     repository.DBRepo
	userRepo repository.UserRepository
	service  *service.UserService
}

// NewUserUsecase a user usecase package
func NewUserUsecase(repo repository.DBRepo, userRepo repository.UserRepository, service *service.UserService) UserUsecase {
	return &userUsecase{
		repo:     repo,
		userRepo: userRepo,
		service:  service,
	}
}

func (u *userUsecase) ListUser() ([]model.User, error) {
	users, err := u.userRepo.FindAll(u.repo, repository.QueryParam{})
	if err != nil {
		return nil, err
	}

	rs := []model.User{}
	for idx := range users {
		if users[idx].Name == "Ân" {
			continue
		}
		rs = append(rs, users[idx])
	}
	return rs, nil
}

func (u *userUsecase) RegisterUser(input model.User) (*model.User, error) {
	newRepo, finallyFn := u.repo.NewTransaction()

	var err error

	defer func() {
		finallyFn(err)
	}()

	if err = u.service.Duplicated(newRepo, input.Email); err != nil {
		return nil, err
	}

	if err := u.userRepo.Save(u.repo, &input); err != nil {
		return nil, err
	}

	// TODO:
	// if err := u.ThirdParty.Validate(user); err != nil {
	// 	Rollback()
	// 	return nil, err
	// }

	return &input, nil
}
