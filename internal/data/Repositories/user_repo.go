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

// Create Account
// Log in
// validate session
// withdraw session (if stateful)
// log out (if stateful, or just expire token)
// Refresh token (SHould token be one use only?)
