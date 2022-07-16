package usecase

import (
	"errors"
	"testing"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSession(t *testing.T) {
	testTable := []struct {
		name, Name, Date, Start, End             		string
		update                                  func(id int, rent model.Session) error
		noError			                        bool
		//status									*bool
		id, LocationID, VaccineID				int
		//StockVaccine							*int64
	}{
		{
			name: "success",
			update: func(id int, rent model.Session) error {
				return nil
			},
			noError:         true,
			id: 				1,
			LocationID: 		1,
			VaccineID: 			1,
			Name: 				"moderna",
			Date: 				"2022-16-07 00:00:00",
			Start: 				"2022-16-07 00:00:00",
			End: 				"2022-17-07 00:00:00",
			//status: true,
			//StockVaccine: 12,
		},
		{
			name: "error internal",
			update: func(id int, rent model.Session) error {
				return errors.New("error")
			},
			noError:         false,
			id: 				2,
			LocationID: 		1,
			VaccineID: 			1,
			Name: 				"moderna",
			Date: 				"2022-16-07 00:00:00",
			Start: 				"2022-16-07 00:00:00",
			End: 				"2022-17-07 00:00:00",
			//status:          *false,
			//StockVaccine: 24,
		},
	}
	repoSession := RepoMockSession{}
	
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repoSession.update = v.update

			svc := NewSession(&repoSession, config.Config{})
			err := svc.UpdateSession(v.id, model.Session{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteSession(t *testing.T) {
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
	repoSession := RepoMockSession{}


	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repoSession.delete = v.delete

			svc := NewSession(&repoSession, config.Config{})
			err := svc.DeleteSessionByID(v.id)
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
