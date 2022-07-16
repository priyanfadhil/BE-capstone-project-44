package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterUserRepository interface {
	CreateUsers(user model.User) (model.User, error)
	GetAllUsers() []model.User
	GetOneUserByID(id int) (user model.User, err error)
	GetOneUserByEmail(email string) (user model.User, err error)
	UpdateOneUserByID(id int, user model.User) error
	DeleteUserByID(id int) error
}

// Use Case
type AdapterUser interface {
	CreateUser(email string, user model.User) error
	(id, idToken int, user mUpdateUserodel.User) error
	GetAllUsers() []model.User
	GetUserByID(id int) (model.User, error)
	LoginUser(email, password string) (string, int)
	DeleteUserByID(id int) error
}
