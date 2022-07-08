package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type AdminRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *AdminRepositoryMysqlLayer) CreateAdmins(admin model.Admin) error {
	res := r.DB.Create(&admin)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *AdminRepositoryMysqlLayer) GetAllAdmins() []model.Admin {
	vanncine := []model.Admin{}
	r.DB.Find(&vanncine)

	return vanncine
}

func (r *AdminRepositoryMysqlLayer) GetOneAdminByID(id int) (admin model.Admin, err error) {
	res := r.DB.Where("id = ?", id).Find(&admin)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *AdminRepositoryMysqlLayer) GetOneAdminByEmail(email string) (admin model.Admin, err error) {
	res := r.DB.Where("email = ?", email).Find(&admin)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return admin, err
}

func (r *AdminRepositoryMysqlLayer) UpdateOneAdminByID(id int, admin model.Admin) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&admin)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *AdminRepositoryMysqlLayer) DeleteAdminByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Admin{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewAdminMysqlRepository(db *gorm.DB) domain.AdapterAdminRepository {
	return &AdminRepositoryMysqlLayer{
		DB: db,
	}
}
