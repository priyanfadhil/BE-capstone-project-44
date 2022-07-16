package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterAdminRepository interface {
	CreateAdmins(admin model.Admin) error
	GetAllAdmins() []model.Admin
	GetOneAdminByID(id int) (admin model.Admin, err error)
	GetOneAdminByEmail(email string) (user model.Admin, err error)
	UpdateOneAdminByID(id int, admin model.Admin) error
	DeleteAdminByID(id int) error
}

// Use Case
type AdapterAdmin interface {
	CreateAdmin(email string, admin model.Admin) error
	UpdateAdmin(id int, admin model.Admin) error
	GetAllAdmins() []model.Admin
	GetAdminByID(id int) (model.Admin, error)
	LoginAdmin(email, password string) (string, int)
	DeleteAdminByID(id int) error
}
