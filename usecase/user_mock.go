package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockUser struct {
	update func(id int, user model.User) error
	create func(user model.User) (model.User, error)
	delete func(id int) error
	getone func(id int) (user model.User, err error)
}

func (r *RepoMockUser) CreateUsers(user model.User) (model.User, error) {
	return r.create(user)
}
func (r *RepoMockUser) GetAllUsers() []model.User {
	panic("implemente")
}
func (r *RepoMockUser) GetOneUserByID(id int) (user model.User, err error) {
	return r.getone(id)
}
func (r *RepoMockUser) UpdateOneUserByID(id int, user model.User) error {
	return r.update(id, user)
}
func (r *RepoMockUser) DeleteUserByID(id int) error {
	return r.delete(id)
}

func (r *RepoMockUser) GetOneUserByEmail(email string) (user model.User, err error) {
	panic("implemente")
}
