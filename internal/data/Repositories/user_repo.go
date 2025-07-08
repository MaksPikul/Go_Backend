package repositories

import "database/sql"

type UserRepositoryInterface interface {
	//Create(*models.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
