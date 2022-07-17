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

func TestGetOneFamilyMembers(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := FamilyMemberMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookings` WHERE `bookings`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "nik", "relationship", "name", "birthdate", "gender", "phone", "address"}).
			AddRow(1, 1, "1234567890123456", "anak", "ikhsan", "2000-26-05", "pr", "08939495169", "jln. maguwo"))

	res := repo.GetAllFamilyMembers()
	assert.Equal(t, res[0].Name, "ikhsan")
	assert.Len(t, res, 1)
}

func TestGetAllFamilyMembers(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := FamilyMemberMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `bookings`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "nik", "relationship", "name", "birthdate", "gender", "phone", "address"}).
			AddRow(1, 1, "1234567890123456", "anak", "ikhsan", "2000-26-05", "pr", "08939495169", "jln. maguwo").
			AddRow(2, 2, "1234567890123430", "ayah", "sule", "1990-26-05", "lk", "08939495179", "jln. maguwo").
			AddRow(3, 3, "1234567890123450", "ibu", "mawar", "1990-26-05", "pr", "08939495168", "jln. maguwo"))

	res := repo.GetAllFamilyMembers()
	assert.Equal(t, res[0].Name, "ikhsan")
	assert.Len(t, res, 3)
}

func TestDeleteFamilyMemberByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := FamilyMemberMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteFamilyMemberByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateFamilyMemberByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := FamilyMemberMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs("1234567890123456", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneFamilyMemberByID(1, model.FamilyMember{
		NIK: "1234567890123456",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}
