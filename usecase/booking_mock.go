package usecase

import "github.com/priyanfadhil/BE-capstone-project-44/model"

type RepoMockBooking struct {
	update func(id int, booking model.Booking) error
	create func(booking model.Booking) error
	delete func(id int) error
	getone func(id int) (booking model.Booking, err error)
}

func (r *RepoMockBooking) CreateBooking(booking model.Booking) error {
	return r.create(booking)
}
func (r *RepoMockBooking) GetAllBookings() []model.Booking {
	panic("implemente")
}
func (r *RepoMockBooking) GetOneBookingByID(id int) (booking model.Booking, err error) {
	return r.getone(id)
}
func (r *RepoMockBooking) UpdateOneBookingByID(id int, booking model.Booking) error {
	return r.update(id, booking)
}
func (r *RepoMockBooking) DeleteBookingByID(id int) error {
	return r.delete(id)
}
