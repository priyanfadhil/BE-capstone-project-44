package usecase

import (
	"errors"
	"testing"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateVaccinationLocation(t *testing.T) {
	testTable := []struct {
		name, Name, Address                             string
		update                                  		func(id int, rent model.VaccinationLocation) error
		noError                                 		bool
		id, CreatedBY 									int
	}{
		{
			name: "success",
			update: func(id int, rent model.VaccinationLocation) error {
				return nil
			},
			noError:         	true,
			id:              	1,
			CreatedBY: 			1,
			Name:        		"moderna",
			Address: 		"Jl. Pasteur No.38, Pasteur, Kec. Sukajadi, Kota Bandung, Jawa Barat 40161",

		},
		{
			name: "error internal",
			update: func(id int, rent model.VaccinationLocation) error {
				return errors.New("error")
			},
			noError:         false,
			id:              	2,
			CreatedBY: 			2,
			Name:        		"moderna",
			Address: 			"Jl. maguwo No.38, Kec. , Kota Yogyakarta, DIY 40161",

		},
	}
	repoVaccinationLocation := RepoMockVaccinationLocation{}
	
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repoVaccinationLocation.update = v.update

			svc := NewVaccinationLocation(&repoVaccinationLocation, config.Config{})
			err := svc.UpdateVaccinationLocation(v.id, model.VaccinationLocation{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteVaccinationLocation(t *testing.T) {
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
	repoVaccinationLocation := RepoMockVaccinationLocation{}
	

	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repoVaccinationLocation.delete = v.delete

			svc := NewVaccinationLocation(&repoVaccinationLocation, config.Config{})
			err := svc.DeleteVaccinationLocationByID(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

// func TestGetOneVaccine(t *testing.T) {
// 	book := model.Vaccineing{
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
