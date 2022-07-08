package usecase

import (
	"fmt"
	"net/http"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/middleware"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcUser struct {
	c    config.Config
	repo domain.AdapterUserRepository
}

func (s *svcUser) CreateUser(email string, user model.User) error {
	newUser, _ := s.repo.GetOneUserByEmail(email)

	if newUser.Email == email {
		return fmt.Errorf("email sudah terdaftar")
	}

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

func (s *svcUser) LoginUser(name, password string) (string, int) {
	user, _ := s.repo.GetOneUserByEmail(name)
	print(user.Password)

	if user.Password != password {
		return "", http.StatusUnauthorized
	}

	token, err := middleware.CreateToken(user.ID, user.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func User(repo domain.AdapterUserRepository, c config.Config) domain.AdapterUser {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
