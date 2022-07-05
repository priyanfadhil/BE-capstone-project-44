package usecase

import (
	"fmt"
	"net/http"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/middleware"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcAdmin struct {
	c    config.Config
	repo domain.AdapterAdminRepository
}

func (s *svcAdmin) CreateAdmin(admin model.Admin) error {
	return s.repo.CreateAdmins(admin)
}

func (s *svcAdmin) UpdateAdmin(id int, admin model.Admin) error {
	fmt.Println(id)
	return s.repo.UpdateOneAdminByID(id, admin)
}

func (s *svcAdmin) GetAllAdmins() []model.Admin {
	return s.repo.GetAllAdmins()
}

func (s *svcAdmin) GetAdminByID(id int) (model.Admin, error) {
	return s.repo.GetOneAdminByID(id)
}

func (s *svcAdmin) DeleteAdminByID(id int) error {
	return s.repo.DeleteAdminByID(id)
}

func (s *svcAdmin) LoginAdmin(email, password string) (string, int) {
	admin, _ := s.repo.GetOneAdminByEmail(email)
	print(admin.Password)

	if admin.Password != password {
		return "", http.StatusUnauthorized
	}

	token, err := middleware.CreateToken(admin.ID, admin.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func NewAdmin(repo domain.AdapterAdminRepository, c config.Config) domain.AdapterAdmin {
	return &svcAdmin{
		repo: repo,
		c:    c,
	}
}
