package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcBooking struct {
	c    config.Config
	repo domain.AdapterBookingRepository
}

func (s *svcBooking) CreateBooking(booking model.Booking) error {
	return s.repo.CreateBooking(booking)
}

func (s *svcBooking) UpdateBooking(id int, booking model.Booking) error {
	fmt.Println(id)
	return s.repo.UpdateOneBookingByID(id, booking)
}

func (s *svcBooking) GetAllBookings() []model.Booking {
	return s.repo.GetAllBookings()
}

func (s *svcBooking) GetBookingByID(id int) (model.Booking, error) {
	return s.repo.GetOneBookingByID(id)
}

func (s *svcBooking) DeleteBookingByID(id int) error {
	return s.repo.DeleteBookingByID(id)
}

func NewBooking(repo domain.AdapterBookingRepository, c config.Config) domain.AdapterBooking {
	return &svcBooking{
		repo: repo,
		c:    c,
	}
}
