package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockUserAddress struct {
	update func(id int, useraddress model.UserAddress) error
	create func(useraddress model.UserAddress) error
	delete func(id int) error
	getone func(id int) (useraddress model.UserAddress, err error)
}

func (r *RepoMockUserAddress) CreateUserAddress(useraddress model.UserAddress) error {
	return r.create(useraddress)
}
func (r *RepoMockUserAddress) GetAllUserAddresses() []model.UserAddress {
	panic("implemente")
}
func (r *RepoMockUserAddress) GetOneUserAddressByID(id int) (useraddress model.UserAddress, err error) {
	return r.getone(id)
}
func (r *RepoMockUserAddress) UpdateOneUserAddressByID(id int, useraddress model.UserAddress) error {
	return r.update(id, useraddress)
}
func (r *RepoMockUserAddress) DeleteUserAddressByID(id int) error {
	return r.delete(id)
}
