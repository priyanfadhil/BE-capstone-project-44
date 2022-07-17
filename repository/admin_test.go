package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetOneAdmins(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewAdminMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `admins` WHERE `admins`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "role"}).
			AddRow(1, "mia@gmail.com", "Mia123.", "super admin"))

	res := repo.GetAllAdmins()
	assert.Equal(t, res[0].Email, "mia@gmail.com")
	assert.Len(t, res, 1)
}

func TestGetAllAdmins(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewAdminMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `admins`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "role"}).
			AddRow(1, "mia@gmail.com", "Mia123.", "super admin").
			AddRow(2, "raju@gmail.com", "raju123.", "super admin").
			AddRow(3, "vanya@gmail.com", "vanya123.", "super admin"))

	res := repo.GetAllAdmins()
	assert.Equal(t, res[0].Email, "mia@gmail.com")
	assert.Len(t, res, 3)
}

func TestDeleteAdminByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewAdminMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteAdminByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateAdminByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewAdminMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("mia@gmail.com", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneAdminByID(1, model.Admin{
		Email: "mia@gmail.com",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
