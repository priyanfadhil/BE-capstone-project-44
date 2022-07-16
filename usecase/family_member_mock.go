package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockFamilyMember struct {
	update func(id int, familymember model.FamilyMember) error
	create func(familymember model.FamilyMember) error
	delete func(id int) error
	getone func(id int) (familymember model.FamilyMember, err error)
}

func (r *RepoMockFamilyMember) CreateFamilyMembers(familymember model.FamilyMember) error {
	return r.create(familymember)
}
func (r *RepoMockFamilyMember) GetAllFamilyMembers() []model.FamilyMember {
	panic("implemente")
}
func (r *RepoMockFamilyMember) GetOneFamilyMemberByID(id int) (familymember model.FamilyMember, err error) {
	return r.getone(id)
}
func (r *RepoMockFamilyMember) UpdateOneFamilyMemberByID(id int, familymember model.FamilyMember) error {
	return r.update(id, familymember)
}
func (r *RepoMockFamilyMember) DeleteFamilyMemberByID(id int) error {
	return r.delete(id)
}




