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

func TestGetVaccineStatus(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewVaccineStatusMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `Vaccine_Statuses` WHERE `Vaccine_Statuses`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status", "booking"}).
			AddRow(1, "vaccine pertama" ))

	res := repo.GetAllVaccineStatuses()
	assert.Equal(t, res[0].Status, "vaccine pertama")
	assert.Len(t, res, 1)
}

func TestGetAllVaccineStatuses(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewVaccineStatusMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `Vaccine_Statuses`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status"}).
			AddRow(1, "vaccine pertama").
			AddRow(2, "vaccine kedua").
			AddRow(3, "vaccine ketiga"))

	res := repo.GetAllVaccineStatuses()
	assert.Equal(t, res[0].Status, "vaccine pertama")
	assert.Len(t, res, 3)
}

func TestDeleteVaccineStatusyID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewVaccineStatusMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteVaccineStatusByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateVaccineStatusyID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewVaccineStatusMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(true, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneVaccineStatusByID(1, model.VaccineStatus{
		Status: "vaccine pertama",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
