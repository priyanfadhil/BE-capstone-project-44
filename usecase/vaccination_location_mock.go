package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockVaccinationLocation struct {
	update func(id int, vaccinationlocation model.VaccinationLocation) error
	create func(vaccinationlocation model.VaccinationLocation) error
	delete func(id int) error
	getone func(id int) (vaccinationlocation model.VaccinationLocation, err error)
}

func (r *RepoMockVaccinationLocation) CreateVaccinationLocation(vaccinationlocation model.VaccinationLocation) error {
	return r.create(vaccinationlocation)
}
func (r *RepoMockVaccinationLocation) GetAllVaccinationLocations() []model.VaccinationLocation {
	panic("implemente")
}
func (r *RepoMockVaccinationLocation) GetOneVaccinationLocationByID(id int) (VaccinationLocation model.VaccinationLocation, err error) {
	return r.getone(id)
}
func (r *RepoMockVaccinationLocation) UpdateOneVaccinationLocationByID(id int, VaccinationLocation model.VaccinationLocation) error {
	return r.update(id, vaccinationlocation)
}
func (r *RepoMockVaccinationLocation) DeleteVaccinationLocationByID(id int) error {
	return r.delete(id)
}
