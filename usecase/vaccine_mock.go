package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockVaccine struct {
	update func(id int, vaccine model.Vaccine) error
	create func(vaccine model.Vaccine) error
	delete func(id int) error
	getone func(id int) (vaccine model.Vaccine, err error)
}

func (r *RepoMockVaccine) CreateVaccine(vaccine model.Vaccine) error {
	return r.create(vaccine)
}
func (r *RepoMockVaccine) GetAllVaccines() []model.Vaccine {
	panic("implemente")
}
func (r *RepoMockVaccine) GetOneVaccineByID(id int) (vaccine model.Vaccine, err error) {
	return r.getone(id)
}
func (r *RepoMockVaccine) UpdateOneVaccineByID(id int, vaccine model.Vaccine) error {
	return r.update(id, vaccine)
}
func (r *RepoMockVaccine) DeleteVaccineByID(id int) error {
	return r.delete(id)
}
