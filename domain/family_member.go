package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterFamilyMemberRepository interface {
	CreateFamilyMembers(familymember model.FamilyMember) error
	GetAllFamilyMembers() []model.FamilyMember
	GetOneFamilyMemberByID(id int) (familymember model.FamilyMember, err error)
	UpdateOneFamilyMemberByID(id int, familymember model.FamilyMember) error
	DeleteFamilyMemberByID(id int) error
}

// Use Case
type AdapterFamilyMember interface {
	CreateFamilyMember(familymember model.FamilyMember) error
	UpdateFamilyMember(id, idToken int, familymember model.FamilyMember) error
	GetAllFamilyMembers() []model.FamilyMember
	GetFamilyMemberByID(id int) (model.FamilyMember, error)
	DeleteFamilyMemberByID(id int) error
}
