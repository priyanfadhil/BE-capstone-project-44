package domain

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

// repository
type AdapterBookingRepository interface {
	CreateBooking(Booking model.Booking) error
	GetAllBookings() []model.Booking
	GetOneBookingByID(id int) (Booking model.Booking, err error)
	UpdateOneBookingByID(id int, Booking model.Booking) error
	DeleteBookingByID(id int) error
}

// Use Case
type AdapterBooking interface {
	CreateBooking(Booking model.Booking) error
	UpdateBooking(id int, Booking model.Booking) error
	GetAllBookings() []model.Booking
	GetBookingByID(id int) (model.Booking, error)
	DeleteBookingByID(id int) error
}
