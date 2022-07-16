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

func TestGetOneVaccines(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewVaccineMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vaccines` WHERE `vaccines`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "moderna"))

	res := repo.GetAllVaccines()
	assert.Equal(t, res[0].Name,"moderna")
	assert.Len(t, res, 1)
}

func TestGetAllVaccines(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewVaccineMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vaccines`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "moderna").
			AddRow(2, "moderna").
			AddRow(3, "moderna"))

	res := repo.GetAllVaccines()
	assert.Equal(t, res[0].Name,"moderna")
	assert.Len(t, res, 3)
}

func TestDeleteVaccineByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewVaccineMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteVaccineByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateVaccineByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewVaccineMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(true, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneVaccineByID(1, model.Vaccine{
		Name: "moderna",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
