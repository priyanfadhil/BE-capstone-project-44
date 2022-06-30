package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type UserAddressRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *UserAddressRepositoryMysqlLayer) CreateUserAddress(useraddress model.UserAddress) error {
	res := r.DB.Create(&useraddress)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *UserAddressRepositoryMysqlLayer) GetAllUserAddresses() []model.UserAddress {
	useraddress := []model.UserAddress{}
	r.DB.Find(&useraddress)
	//r.DB.Model(&model.User{}).Association("rents").Find(&users)

	return useraddress
}

func (r *UserAddressRepositoryMysqlLayer) GetOneUserAddressByID(id int) (useraddress model.UserAddress, err error) {
	res := r.DB.Where("id = ?", id).Find(&useraddress)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *UserAddressRepositoryMysqlLayer) UpdateOneUserAddressByID(id int, useraddress model.UserAddress) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&useraddress)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *UserAddressRepositoryMysqlLayer) DeleteUserAddressByID(id int) error {
	res := r.DB.Delete(&model.UserAddress{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewUserAddressMysqlRepository(db *gorm.DB) domain.AdapterUserAddressRepository {
	return &UserAddressRepositoryMysqlLayer{
		DB: db,
	}
}
