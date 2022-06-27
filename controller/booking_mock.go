package controller

import (
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"github.com/stretchr/testify/mock"
)

type MockBookingSvc struct {
	Mock mock.Mock
}

func (m *MockBookingSvc) CreateBooking(session int, booking model.Booking) error {
	arg := m.Mock.Called(session, booking)

	// if arg.Get(0) == nil {
	// 	return nil
	// } else {
	// 	return arg.Error(0)
	// }
	// //return ret.Error(0)

	return arg.Error(0)
}
func (m *MockBookingSvc) UpdateBooking(id int, booking model.Booking) error {
	return nil
}
func (m *MockBookingSvc) GetAllBookings() []model.Booking {
	return []model.Booking{}
}
func (m *MockBookingSvc) GetBookingByID(id int) (model.Booking, error) {
	return model.Booking{}, nil
}
func (m *MockBookingSvc) DeleteBookingByID(id int) error {
	return nil
}
