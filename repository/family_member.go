package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type familymemberRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *familymemberRepositoryMysqlLayer) CreateFamilyMembers(familymember model.FamilyMember) error {
	res := r.DB.Create(&familymember)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *familymemberRepositoryMysqlLayer) GetAllFamilyMembers() []model.FamilyMember {
	familymembers := []model.FamilyMember{}
	r.DB.Find(&familymembers)

	return familymembers
}

func (r *familymemberRepositoryMysqlLayer) GetOneFamilyMemberByID(id int) (familymember model.FamilyMember, err error) {
	res := r.DB.Where("id = ?", id).Find(&familymember)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *familymemberRepositoryMysqlLayer) UpdateOneFamilyMemberByID(id int, familymember model.FamilyMember) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&familymember)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *familymemberRepositoryMysqlLayer) DeleteFamilyMemberByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.FamilyMember{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func FamilyMemberMysqlRepository(db *gorm.DB) domain.AdapterFamilyMemberRepository {
	return &familymemberRepositoryMysqlLayer{
		DB: db,
	}
}
