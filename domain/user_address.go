package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterUserAddressRepository interface {
	CreateUserAddress(UserAddress model.UserAddress) error
	GetAllUserAddress() []model.UserAddress
	GetOneUserAddressByID(id int) (UserAddress model.UserAddress, err error)
	UpdateOneUserAddressByID(id int, UserAddress model.UserAddress) error
	DeleteUserAddressByID(id int) error
}

// Use Case
type AdapterUserAddress interface {
	CreateUserAddress(UserAddress model.UserAddress) error
	UpdateUserAddress(id int, UserAddress model.UserAddress) error
	GetAllUserAddress() []model.UserAddress
	GetUserAddressByID(id int) (model.UserAddress, error)
	DeleteUserAddressByID(id int) error
}
