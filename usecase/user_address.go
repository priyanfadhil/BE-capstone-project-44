package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcUserAddress struct {
	c    config.Config
	repo domain.AdapterUserAddressRepository
}

func (s *svcUserAddress) CreateUserAddress(useraddress model.UserAddress) error {
	return s.repo.CreateUserAddress(useraddress)
}

func (s *svcUserAddress) UpdateUserAddress(id int, useraddress model.UserAddress) error {
	fmt.Println(id)
	return s.repo.UpdateOneUserAddressByID(id, useraddress)
}

func (s *svcUserAddress) GetAllUserAddress() []model.UserAddress {
	return s.repo.GetAllUserAddress()
}

func (s *svcUserAddress) GetUserAddressByID(id int) (model.UserAddress, error) {
	return s.repo.GetOneUserAddressByID(id)
}

func (s *svcUserAddress) DeleteUserAddressByID(id int) error {
	return s.repo.DeleteUserAddressByID(id)
}

func NewUserAddress(repo domain.AdapterUserAddressRepository, c config.Config) domain.AdapterUserAddress {
	return &svcUserAddress{
		repo: repo,
		c:    c,
	}
}
