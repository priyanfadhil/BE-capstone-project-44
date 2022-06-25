package usecase

import (
	"fmt"
	"net/http"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/helper"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcVaccineStatus struct {
	c    config.Config
	repo domain.AdapterVaccineStatusRepository
}

func (s *svcVaccineStatus) CreateVaccineStatus(vaccinestatus model.VaccineStatus) error {
	return s.repo.CreateVaccine(vaccine)
}

func (s *svcVaccineStatus) UpdateVaccineStatus(id, idToken int, vaccinestatus model.VaccineStatus) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneVaccineStatusByID(id, vaccinestatus)
}

func (s *svcVaccineStatus) GetAllVaccineStatus() []model.VaccineStatus {
	return s.repo.GetAllVaccineStatus()
}

func (s *svcVaccineStatus) GetVaccineStatusByID(id int) (model.VaccineStatus, error) {
	return s.repo.GetOneVaccineStatusByID(id)
}

func (s *svcVaccineStatus) DeleteVaccineStatusByID(id int) error {
	return s.repo.DeleteVaccineStatusByID(id)
}

func NewVaccineStatus(repo domain.AdapterVaccineStatusRepository, c config.Config) domain.AdapterVaccineStatus {
	return &svcVaccineStatus{
		repo: repo,
		c:    c,
	}
}

func (s *svcVaccineStatus) LoginVaccineStatus(name, password string) (string, int) {
	vaccinestatus, _ := s.repo.GetOneVaccineStatusByName(name)

	if vaccinestatus.Password != password {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(vaccinestatus.ID, vaccinestatus.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
