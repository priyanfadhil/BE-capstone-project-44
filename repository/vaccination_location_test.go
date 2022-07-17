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

func TestGetOneVaccinationLocations(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := VaccinationLocationMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookings` WHERE `bookings`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_by", "session_id", "name", "address"}).
			AddRow(1, 1, 1, "moderna", "jln. maguwo"))

	res := repo.GetAllVaccinationLocations()
	assert.Equal(t, res[0].Name, "moderna")
	assert.Len(t, res, 1)
}

func TestGetAllVaccinationLocationss(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := VaccinationLocationMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookings`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_by", "session_id", "name", "address"}).
			AddRow(1, 1, 1, "moderna", "jln. maguwo").
			AddRow(2, 2, 2, "astrazeneca", "jln. saturnus").
			AddRow(3, 3, 3, "sinovac", "jln. mawar"))

	res := repo.GetAllVaccinationLocations()
	assert.Equal(t, res[0].Name, "moderna")
	assert.Len(t, res, 3)
}

func TestDeleteVaccinationLocationByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := VaccinationLocationMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteVaccinationLocationByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateVaccinationLocationByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := VaccinationLocationMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("moderna", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneVaccinationLocationByID(1, model.VaccinationLocation{
		Name: "moderna",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
