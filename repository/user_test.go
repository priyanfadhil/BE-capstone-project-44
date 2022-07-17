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

func TestGetOneUser(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := UserMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nik", "name", "email", "phone", "password"}).
			AddRow(1, "1671052285101234", "mia", "mia@gmail.com","087758960520", "lebahganteng"))

	res := repo.GetAllUsers()
	assert.Equal(t, res[0].NIK,"1671052285101234" )
	assert.Len(t, res, 1)
}

func TestGetAllUsers(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := UserMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nik", "name", "email", "phone", "password"}).
			AddRow(1, "1671052285101234", "mia", "mia@gmail.com","087758960520", "lebahganteng").
			AddRow(2, "1671052285101298", "vanya", "vanya@gmail.com","081358960560", "boeing737").
			AddRow(3, "1671052285101297", "lusi", "lusi@gmail.com","085458960590", "grobe"))

	res := repo.GetAllUsers()
	assert.Equal(t, res[0].NIK,"1671052285101234")
	assert.Len(t, res, 3)
}

func TestDeleteUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := UserMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteUserByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateUserByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := UserMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("1671052285101234", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneUserByID(1, model.User{
		NIK: "1671052285101234",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
