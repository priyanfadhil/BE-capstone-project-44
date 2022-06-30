package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockVaccineStatus struct {
	update func(id int, vaccinestatus model.VaccineStatus) error
	create func(vaccinestatus model.VaccineStatus) error
	delete func(id int) error
	getone func(id int) (vaccinestatus model.VaccineStatus, err error)
}

func (r *RepoMockVaccineStatus) CreateVaccineStatus(vaccinestatus model.VaccineStatus) error {
	return r.create(vaccinestatus)
}
func (r *RepoMockVaccineStatus) GetAllVaccinationStatus() []model.VaccineStatus {
	panic("implemente")
}
func (r *RepoMockVaccineStatus) GetOneVaccinationStatusByID(id int) (vaccinestatus model.VaccineStatus, err error) {
	return r.getone(id)
}
func (r *RepoMockVaccineStatus) UpdateOneVaccinationStatusByID(id int, vaccinestatus model.VaccineStatus) error {
	return r.update(id, vaccinestatus)
}
func (r *RepoMockVaccineStatus) DeleteVaccinationStatusByID(id int) error {
	return r.delete(id)
}
