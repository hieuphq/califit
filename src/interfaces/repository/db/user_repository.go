package db

import (
	"github.com/hieuphq/califit/src/domain/model"
	"github.com/hieuphq/califit/src/interfaces/repository"
)

// userRepository implimentation of User Repository
type userRepository struct {
}

// NewUserRepository new a user repository
func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

// FindAll get all user in db
func (us *userRepository) FindAll(repo repository.DBRepo, param repository.QueryParam) ([]model.User, error) {
	db := repo.DB()
	users := []model.User{}

	query := db.Table(model.UserTable)

	if param.Limit > 0 {
		query.Limit(param.Limit).Offset(param.Offset)
	}

	if len(param.OrderConditions) > 0 {
		for idx := range param.OrderConditions {
			o := param.OrderConditions[idx]
			query.Order(o.Field, o.IsAscending)
		}
	}
	return users, query.Find(&users).Error
}

// FindByEmail user with key is email
func (us *userRepository) FindByEmail(repo repository.DBRepo, email string) (*model.User, error) {
	db := repo.DB()
	var user model.User
	query := db.Where("email = ?", email)
	return &user, query.Find(&user).Error
}

// Save user to db
func (us *userRepository) Save(repo repository.DBRepo, user *model.User) error {
	db := repo.DB()

	return db.Model(&model.User{}).
		Where("id = ?", user.ID).
		Save(&user).Error
}
