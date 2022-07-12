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
	c          config.Config
	Userrepo   domain.AdapterUserRepository
	Familyrepo domain.AdapterFamilyMemberRepository
}

func (s *svcUser) CreateUser(email string, user model.User) error {
	newUser, _ := s.Userrepo.GetOneUserByEmail(email)
	family := model.FamilyMember{}

	if newUser.Email == email {
		return fmt.Errorf("email sudah terdaftar")
	}

	create, err := s.Userrepo.CreateUsers(user)

	if err != nil {
		return err
	}

	family.UserID = create.ID
	family.NIK = user.NIK
	family.Name = user.Name
	family.Phone = user.Phone

	s.Familyrepo.CreateFamilyMembers(family)
	return err
}

func (s *svcUser) UpdateUser(id, idToken int, user model.User) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.Userrepo.UpdateOneUserByID(id, user)
}

func (s *svcUser) GetAllUsers() []model.User {
	return s.Userrepo.GetAllUsers()
}

func (s *svcUser) GetUserByID(id int) (model.User, error) {
	return s.Userrepo.GetOneUserByID(id)
}

func (s *svcUser) DeleteUserByID(id int) error {
	return s.Userrepo.DeleteUserByID(id)
}

func (s *svcUser) LoginUser(name, password string) (string, int) {
	user, _ := s.Userrepo.GetOneUserByEmail(name)
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

func User(Userrepo domain.AdapterUserRepository, Familyrepo domain.AdapterFamilyMemberRepository, c config.Config) domain.AdapterUser {
	return &svcUser{
		Userrepo:   Userrepo,
		Familyrepo: Familyrepo,
		c:          c,
	}
}
