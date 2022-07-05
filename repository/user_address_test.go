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

func TestGetOneUserAddress(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewUserAddressMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `UserAddresss` WHERE `UserAddress`.`deleted_at` IS NULL")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "alamat", "kelurahan", "kecamatan", "kota", "provinsi", "vaccinationlocations"}).
			AddRow(1, 1, "jln. maguwo", "banguntapan", "banguntapan", "yogyakarta", "DIY", "rs. hardjolukito"))

	res := repo.GetAllUserAddresses()
	assert.Equal(t, res[0].Alamat, "jln. maguwo")
	assert.Len(t, res, 1)
}

func TestGetAllUserAddresss(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewUserAddressMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `UserAddresss`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "alamat", "kelurahan", "kecamatan", "kota", "provinsi", "vaccinationlocations"}).
			AddRow(1, 1, "jln. maguwo", "banguntapan", "banguntapan", "bantul", "DIY", "rspau hardjolukito").
			AddRow(2, 2, "jln. kronggahan", "trihanggo", "gamping", "sleman", "DIY", "rsup dr. sardjito ").
			AddRow(3, 3, "jln. laksada adisucipto", "demangan", "gondokusuman", "DIY", "rs. siloam"))

	res := repo.GetAllUserAddresses()
	assert.Equal(t, res[0].Alamat, "jln. maguwo")
	assert.Len(t, res, 3)
}

func TestDeleteUserAddressByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewUserAddressMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	err := repo.DeleteUserAddressByID(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestUpdateUserAddressByID(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewUserAddressMysqlRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
		WithArgs(true, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	fMock.ExpectCommit()

	err := repo.UpdateOneUserAddressByID(1, model.UserAddress{
		Alamat: "jln. laksada adisutjipto",
	})
	assert.NoError(t, err)
	assert.True(t, true)
}