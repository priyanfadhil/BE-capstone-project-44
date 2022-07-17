package usecase

import (
	"errors"
	"testing"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateFamilyMember(t *testing.T) {
	testTable := []struct {
		name, NIK, Relationship, Name, Birthdate, Gender, Phone, Address, Booking       string
		update                                   										func(id int, rent model.FamilyMember) error
		noError                                  										bool
		id, UserID 										int
	}{
		{
			name: "success",
			update: func(id int, rent model.FamilyMember) error {
				return nil
			},
			noError:         true,
			id: 1,
			UserID: 1,
			NIK: "1234567890123456",
			Relationship: "anak",
			Name: "ikhsan",
			Birthdate: "2000-26-05",
			Gender: "pr",
			Phone: "08939495169",
			Address: "jln. maguwo",
		},
		{
			name: "error internal",
			update: func(id int, rent model.FamilyMember) error {
				return errors.New("error")
			},
			noError:         false,
			id: 0,
			UserID: 0,
			NIK: "1234567890123430",
			Relationship: "ayah",
			Name: "sule",
			Birthdate: "1990-26-05",
			Gender: "lk",
			Phone: "08939495179",
			Address: "jln. maguwo",
		},
	}
	repoFamilyMember := RepoMockFamilyMember{}


	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repoFamilyMember.update = v.update

			svc := FamilyMember(&repoFamilyMember, config.Config{})
			err := svc.UpdateFamilyMember(v.id, model.FamilyMember{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteFamilyMember(t *testing.T) {
	testTable := []struct {
		name    string
		delete  func(id int) error
		noError bool
		id      int
	}{
		{
			name: "success",
			delete: func(id int) error {
				return nil
			},
			noError: true,
			id:      1,
		},

		{
			name: "error internal",
			delete: func(id int) error {
				return errors.New("error")
			},
			noError: false,
			id:      1,
		},
	}
	repoFamilyMember := RepoMockFamilyMember{}
	

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repoFamilyMember.delete = v.delete

			svc := FamilyMember(&repoFamilyMember, config.Config{})
			err := svc.DeleteFamilyMemberByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

// func TestGetOneFamilyMembering(t *testing.T) {
// 	book := model.Booking{
// 		ID:              1,
// 		FamilyID:        1,
// 		SessionID:       1,
// 		StatusVaccineID: 1,
// 		Status:          true,
// 	}
// 	testTable := []struct {
// 		name    string
// 		getone  func(id int) (booking model.Booking, err error)
// 		noError bool
// 		id      int
// 	}{
// 		{
// 			name: "success",
// 			getone: func(id int) (booking model.Booking, err error) {
// 				return book, nil
// 			},
// 			noError: true,
// 			id:      1,
// 		},

// 		{
// 			name: "error internal",
// 			getone: func(id int) (booking model.Booking, err error) {
// 				return book, errors.New("error")
// 			},
// 			noError: false,
// 			id:      1,
// 		},
// 	}
// 	repoBook := RepoMockBooking{}
// 	repoSession := RepoMockSession{}

// 	for _, v := range testTable {
// 		t.Run(v.name, func(t *testing.T) {
// 			repoBook.getone = v.getone

// 			svc := NewBooking(&repoBook, &repoSession, config.Config{})
// 			err := svc.DeleteBookingByID(v.id)
// 			if v.noError {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestCreateBooking(t *testing.T) {
// 	testTable := []struct {
// 		name    string
// 		create  func(rent model.Booking) error
// 		noError bool
// 	}{
// 		{
// 			name: "success",
// 			create: func(rent model.Booking) error {
// 				return nil
// 			},
// 			noError: true,
// 		},
// 		{
// 			name: "error internal",
// 			create: func(rent model.Booking) error {
// 				return errors.New("error")
// 			},
// 			noError: false,
// 		},
// 	}
// 	repoBook := RepoMockBooking{}
// 	repoSession := RepoMockSession{}

// 	for _, v := range testTable {
// 		t.Run(v.name, func(t *testing.T) {
// 			repoBook.create = v.create

// 			svc := NewBooking(&repoBook, &repoSession, config.Config{})
// 			err := svc.CreateBooking(1, model.Booking{})
// 			if v.noError {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }
