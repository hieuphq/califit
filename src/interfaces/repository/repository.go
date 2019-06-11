package repository

import (
	"fmt"

	"github.com/hieuphq/califit/src/interfaces/config"
	"github.com/jinzhu/gorm"
)

// FinallyFunc function to finish a transaction
type FinallyFunc = func(error)

// DBRepo ..
type DBRepo interface {
	DB() *gorm.DB
	NewTransaction() (DBRepo, FinallyFunc)
}

// store is implimentation of repository
type store struct {
	Database *gorm.DB
}

// DB database connection
func (s *store) DB() *gorm.DB {
	return s.Database
}

// NewTransaction for database connection
func (s *store) NewTransaction() (newRepo DBRepo, finallyFn FinallyFunc) {
	newDB := s.Database.Begin()

	finallyFn = func(err error) {
		if err != nil {
			newDB.Rollback()
			return
		}
		newDB.Commit()
		return
	}

	return &store{Database: newDB}, finallyFn
}

// NewPostgresStore postgres init by gorm
func NewPostgresStore(cfg *config.Config) (DBRepo, func() error) {
	ds := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.SQLUsername, cfg.SQLPassword,
		cfg.SQLHost, cfg.SQLPort, cfg.SQLDBName,
	)
	db, err := gorm.Open("postgres", ds)
	if err != nil {
		panic(err)
	}

	return &store{Database: db}, db.Close
}

// NewStore postgres init by gorm
func NewStore(db *gorm.DB) DBRepo {
	return &store{Database: db}
}

// Paging ..
type Paging struct {
	Limit  int
	Offset int
}

// OrderCondition ..
type OrderCondition struct {
	Field       string
	IsAscending bool
}

// OrderConditions ..
type OrderConditions []OrderCondition

// QueryParam is param for query request
type QueryParam struct {
	Paging
	OrderConditions
}
