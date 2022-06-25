package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcVaccinationLocation struct {
	c    config.Config
	repo domain.AdapterVaccinationLocationRepository
}

func (s *svcVaccinationLocation) CreateVaccinationLocation(vaccinationlocation model.VaccinationLocation) error {
	return s.repo.CreateVaccinationLocation(vaccinationlocation)
}

func (s *svcVaccinationLocation) UpdateVaccinationLocation(id int, vaccinationlocation model.VaccinationLocation) error {
	fmt.Println(id)
	return s.repo.UpdateOneVaccinationLocationByID(id, vaccinationlocation)
}

func (s *svcVaccinationLocation) GetAllVaccinationLocations() []model.VaccinationLocation {
	return s.repo.GetAllVaccinationLocations()
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
