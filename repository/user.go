package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type userRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *userRepositoryMysqlLayer) CreateUsers(user model.User) (model.User, error) {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return user, fmt.Errorf("error insert")
	}

	return user, nil
}

func (r *userRepositoryMysqlLayer) GetAllUsers() []model.User {
	users := []model.User{}
	r.DB.Find(&users)

	return users
}

func (r *userRepositoryMysqlLayer) GetOneUserByID(id int) (user model.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *userRepositoryMysqlLayer) GetOneUserByEmail(email string) (user model.User, err error) {
	res := r.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return user, err
}

func (r *userRepositoryMysqlLayer) UpdateOneUserByID(id int, user model.User) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *userRepositoryMysqlLayer) DeleteUserByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.User{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func UserMysqlRepository(db *gorm.DB) domain.AdapterUserRepository {
	return &userRepositoryMysqlLayer{
		DB: db,
	}
}
