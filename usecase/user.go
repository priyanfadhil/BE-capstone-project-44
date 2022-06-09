package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcUser struct {
	c    config.Config
	repo domain.AdapterUserRepository
}

func (s *svcUser) CreateUser(user model.User) error {
	return s.repo.CreateUsers(user)
}

func (s *svcUser) UpdateUser(id, idToken int, user model.User) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneUserByID(id, user)
}

func (s *svcUser) GetAllUsers() []model.User {
	return s.repo.GetAllUsers()
}

func (s *svcUser) GetUserByID(id int) (model.User, error) {
	return s.repo.GetOneUserByID(id)
}

func (s *svcUser) DeleteUserByID(id int) error {
	return s.repo.DeleteUserByID(id)
}

func NewUser(repo domain.AdapterUserRepository, c config.Config) domain.AdapterUser {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
