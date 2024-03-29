package usecase

import (
	"errors"
	"testing"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateBooking(t *testing.T) {
	testTable := []struct {
		name                                     string
		update                                   func(id int, rent model.Booking) error
		noError                                  bool
		FamilyID, SessionID, id, StatusVaccineID int
	}{
		{
			name: "success",
			update: func(id int, rent model.Booking) error {
				return nil
			},
			noError:         true,
			FamilyID:        1,
			id:              1,
			SessionID:       1,
			StatusVaccineID: 1,
		},
		{
			name: "error internal",
			update: func(id int, rent model.Booking) error {
				return errors.New("error")
			},
			noError:         false,
			FamilyID:        1,
			id:              0,
			SessionID:       1,
			StatusVaccineID: 1,
		},
	}
	repoBook := RepoMockBooking{}
	repoSession := RepoMockSession{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repoBook.update = v.update

			svc := NewBooking(&repoBook, &repoSession, config.Config{})
			err := svc.UpdateBooking(v.id, model.Booking{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteBooking(t *testing.T) {
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
	repoBook := RepoMockBooking{}
	repoSession := RepoMockSession{}

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repoBook.delete = v.delete

			svc := NewBooking(&repoBook, &repoSession, config.Config{})
			err := svc.DeleteBookingByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

// func TestGetOneBooking(t *testing.T) {
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
