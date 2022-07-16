package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcVaccine struct {
	c    config.Config
	repo domain.AdapterVaccineRepository
}

func (s *svcVaccine) CreateVaccine(vaccine model.Vaccine) error {
	return s.repo.CreateVaccine(vaccine)
}

func (s *svcVaccine) UpdateVaccine(id int, vaccine model.Vaccine) error {
	fmt.Println(id)
	return s.repo.UpdateOneVaccineByID(id, vaccine)
}

func (s *svcVaccine) GetAllVaccines() []model.Vaccine {
	return s.repo.GetAllVaccines()
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
