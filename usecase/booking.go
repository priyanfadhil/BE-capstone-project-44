package usecase

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type svcBooking struct {
	c           config.Config
	bookingRepo domain.AdapterBookingRepository
	sessionRepo domain.AdapterSessionRepository
}

func (s *svcBooking) CreateBooking(session_id int, booking model.Booking) error {
	stok, _ := s.sessionRepo.GetOneSessionByID(session_id)

	// fmt.Println(stok.StockVaccine)
	if *stok.StockVaccine == 0 {
		return fmt.Errorf("error")
	}

	stok.Name = ""
	stok.Date = ""
	stok.Start = ""
	stok.End = ""
	*stok.Status = false
	*stok.StockVaccine = *stok.StockVaccine - 1
	// fmt.Println(stok.StockVaccine)
	s.sessionRepo.UpdateOneSessionByID(session_id, stok)

	return s.bookingRepo.CreateBooking(booking)
}

func (s *svcBooking) UpdateBooking(id int, booking model.Booking) error {
	fmt.Println(id)
	return s.bookingRepo.UpdateOneBookingByID(id, booking)
}

func (s *svcBooking) GetAllBookings() []model.Booking {
	return s.bookingRepo.GetAllBookings()
}

func (s *svcBooking) GetBookingByID(id int) (model.Booking, error) {
	return s.bookingRepo.GetOneBookingByID(id)
}

func (s *svcBooking) DeleteBookingByID(id int) error {
	return s.bookingRepo.DeleteBookingByID(id)
}

func NewBooking(bookingRepo domain.AdapterBookingRepository, sessionRepo domain.AdapterSessionRepository, c config.Config) domain.AdapterBooking {
	return &svcBooking{
		bookingRepo: bookingRepo,
		sessionRepo: sessionRepo,
		c:           c,
	}
}
