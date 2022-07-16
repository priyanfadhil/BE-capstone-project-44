package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockAdmin struct {
	update func(id int, admin model.Admin) error
	create func(admin model.Admin) error
	delete func(id int) error
	getone func(id int) (admin model.Admin, err error)
}

func (r *RepoMockAdmin) CreateAdmins(admin model.Admin) error {
	return r.create(admin)
}
func (r *RepoMockAdmin) GetAllAdmins() []model.Admin {
	panic("implemente")
}
func (r *RepoMockAdmin) GetOneAdminByID(id int) (admin model.Admin, err error) {
	return r.getone(id)
}
func (r *RepoMockAdmin) UpdateOneAdminByID(id int, admin model.Admin) error {
	return r.update(id, admin)
}
func (r *RepoMockAdmin) DeleteAdminByID(id int) error {
	return r.delete(id)
}
