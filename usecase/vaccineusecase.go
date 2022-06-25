package usecase

import (
	"fmt"
	"net/http"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/helper"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcVaccine struct {
	c    config.Config
	repo domain.AdapterVaccineRepository
}

func (s *svcVaccine) CreateVaccine(vaccine model.Vaccine) error {
	return s.repo.CreateVaccine(vaccine)
}

func (s *svcVaccine) UpdateVaccine(id, idToken int, vaccine model.Vaccine) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneVaccineByID(id, vaccine)
}

func (s *svcVaccine) GetAllVaccine() []model.Vaccine {
	return s.repo.GetAllVaccine()
}

func (s *svcVaccine) GetVaccineByID(id int) (model.Vaccine, error) {
	return s.repo.GetOneVaccineByID(id)
}

func (s *svcVaccine) DeleteVaccineByID(id int) error {
	return s.repo.DeleteVaccineByID(id)
}

func NewVaccine(repo domain.AdapterVaccineRepository, c config.Config) domain.AdapterVaccine {
	return &svcVaccine{
		repo: repo,
		c:    c,
	}
}

func (s *svcVaccine) LoginVaccine(name, password string) (string, int) {
	vaccine, _ := s.repo.GetOneVaccineByName(name)

	if vaccine.Password != password {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(vaccine.ID, vaccine.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
