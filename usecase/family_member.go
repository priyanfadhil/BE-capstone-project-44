package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcFamilyMember struct {
	c    config.Config
	repo domain.AdapterFamilyMemberRepository
}

func (s *svcFamilyMember) CreateFamilyMember(familymember model.FamilyMember) error {
	return s.repo.CreateFamilyMembers(familymember)
}

func (s *svcFamilyMember) UpdateFamilyMember(id int, familymember model.FamilyMember) error {
	fmt.Println(id)
	return s.repo.UpdateOneFamilyMemberByID(id, familymember)
}

func (s *svcFamilyMember) GetAllFamilyMembers() []model.FamilyMember {
	return s.repo.GetAllFamilyMembers()
}

func (s *svcFamilyMember) GetFamilyMemberByID(id int) (model.FamilyMember, error) {
	return s.repo.GetOneFamilyMemberByID(id)
}

func (s *svcFamilyMember) DeleteFamilyMemberByID(id int) error {
	return s.repo.DeleteFamilyMemberByID(id)
}

func FamilyMember(repo domain.AdapterFamilyMemberRepository, c config.Config) domain.AdapterFamilyMember {
	return &svcFamilyMember{
		repo: repo,
		c:    c,
	}
}
