package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcVaccineStatus struct {
	c    config.Config
	repo domain.AdapterVaccineStatusRepository
}

func (s *svcVaccineStatus) CreateVaccineStatus(vaccinestatus model.VaccineStatus) error {
	return s.repo.CreateVaccineStatus(vaccinestatus)
}

func (s *svcVaccineStatus) UpdateVaccineStatus(id int, vaccinestatus model.VaccineStatus) error {
	fmt.Println(id)
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
