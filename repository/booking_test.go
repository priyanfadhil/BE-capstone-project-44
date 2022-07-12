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

func TestGetOneBookings(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewBookingMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookings` WHERE `bookings`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "family_id", "session_id", "status_vaccine_id", "status"}).
			AddRow(1, 1, 1, 1, true))

	res := repo.GetAllBookings()
	assert.Equal(t, res[0].FamilyID, 1)
	assert.Len(t, res, 1)
}

func TestGetAllBookings(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewBookingMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookings`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "family_id", "session_id", "status_vaccine_id", "status"}).
			AddRow(1, 1, 1, 1, true).
			AddRow(2, 2, 1, 1, true).
			AddRow(3, 3, 1, 1, true))

	res := repo.GetAllBookings()
	assert.Equal(t, res[0].FamilyID, 1)
	assert.Len(t, res, 3)
}

func TestDeleteBookingByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewBookingMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteBookingByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateBookingByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewBookingMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(true, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneBookingByID(1, model.Booking{
		Done: true,
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
