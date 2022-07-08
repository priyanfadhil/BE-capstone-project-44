package usecase

import (
	"fmt"
	"net/http"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/helper"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcVaccinationLocation struct {
	c    config.Config
	repo domain.AdapterVaccinationLocationRepository
}

func (s *svcVaccinationLocation) CreateVaccinationLocation(vaccinationlocation model.VaccinationLocation) error {
	return s.repo.CreateVaccine(vaccine)
}

func (s *svcVaccinationLocation) UpdateVaccinationLocation(id, idToken int, vaccinationlocation model.VaccinationLocation) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneVaccinationLocationByID(id, vaccinationlocation)
}

func (s *svcVaccinationLocation) GetAllVaccinationLocation() []model.VaccinationLocation {
	return s.repo.GetAllVaccinationLocation()
}

func (s *svcVaccinationLocation) GetVaccinationLocationByID(id int) (model.VaccinationLocation, error) {
	return s.repo.GetOneVaccinationLocationByID(id)
}

func (s *svcVaccinationLocation) DeleteVaccinationLocationByID(id int) error {
	return s.repo.DeleteVaccinationLocationByID(id)
}

func NewVaccinationLocation(repo domain.AdapterVaccinationLocationRepository, c config.Config) domain.AdapterVaccinationLocation {
	return &svcVaccinationLocation{
		repo: repo,
		c:    c,
	}
}

func (s *svcVaccinationLocation) LoginVaccinationLocation(name, password string) (string, int) {
	vaccinationlocation, _ := s.repo.GetOneVaccinationLocationByName(name)

	if vaccinationlocation.Password != password {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(vaccinationlocation.ID, vaccinationlocation.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
