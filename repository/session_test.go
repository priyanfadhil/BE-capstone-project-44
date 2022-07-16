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

func TestGetOneSessions(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewSessionMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sessions` WHERE `sessions`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "location_id", "vaccine_id", "name", "date", "start", "end", "status", "stock_vaccine"}).
			AddRow(1, 1, 1, 1, true))

	res := repo.GetAllSessions()
	assert.Equal(t, res[0].FamilyID, 1)
	assert.Len(t, res, 1)
}

func TestGetAllSessions(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewSessionMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `sessions`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "family_id", "session_id", "status_vaccine_id", "status"}).
			AddRow(1, 1, 1, 1, true).
			AddRow(2, 2, 1, 1, true).
			AddRow(3, 3, 1, 1, true))

	res := repo.GetAllSessions()
	assert.Equal(t, res[0].FamilyID, 1)
	assert.Len(t, res, 3)
}

func TestDeleteSessionByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewSessionMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteSessionByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateSessionByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewSessionMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(true, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneSessionByID(1, model.Session{
		Done: true,
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
