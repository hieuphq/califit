package registry

import (
	"github.com/hieuphq/califit/src/interfaces/repository"
	"github.com/hieuphq/califit/src/interfaces/repository/db"
	"github.com/hieuphq/califit/src/interfaces/repository/service"
	"github.com/hieuphq/califit/src/usecase"
)

// Container ..
type Container interface {
	UserUC() usecase.UserUsecase
}

type implContainer struct {
	userUC usecase.UserUsecase
}

// NewDefaultContainer ..
func NewDefaultContainer(repo repository.DBRepo) Container {
	userRepo := db.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userUC := usecase.NewUserUsecase(repo, userRepo, userService)
	return &implContainer{
		userUC: userUC,
	}
}

// UserUsecase ..
func (c *implContainer) UserUC() usecase.UserUsecase {
	return c.userUC
}
