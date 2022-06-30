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

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `VaccineStatus` WHERE `VaccineStatus`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status", "booking"}).
			AddRow(1, "OK"))

	res := repo.GetAllVaccineStatuses()
	assert.Equal(t, res[0].Status, 1)
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

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `VaccineStatus`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status", "booking"}).
			AddRow(1, 1, 1, 1, true).
			AddRow(2, 2, 1, 1, true).
			AddRow(3, 3, 1, 1, true))

	res := repo.GetAllVaccineStatuses()
	assert.Equal(t, res[0].Booking.FamilyID, 1)
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

	err := repo.UpdateOneVaccineStatusByID(1, model.VaccineStatus{})

	assert.NoError(t, err)
	assert.True(t, true)
}
