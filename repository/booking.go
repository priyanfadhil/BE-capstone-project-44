package repository

import (
	"fmt"

	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
	"gorm.io/gorm"
)

type BookingRepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *BookingRepositoryMysqlLayer) CreateBooking(booking model.Booking) error {
	res := r.DB.Create(&booking)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}
	return nil
}

func (r *BookingRepositoryMysqlLayer) GetAllBookings() []model.Booking {
	bookings := []model.Booking{}
	r.DB.Find(&bookings)
	return bookings
}

func (r *BookingRepositoryMysqlLayer) GetOneBookingByID(id int) (booking model.Booking, err error) {
	res := r.DB.Where("id = ?", id).Find(&booking)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}
	return
}

func (r *BookingRepositoryMysqlLayer) UpdateOneBookingByID(id int, booking model.Booking) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&booking)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}
	return nil
}

func (r *BookingRepositoryMysqlLayer) DeleteBookingByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Booking{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewBookingMysqlRepository(db *gorm.DB) domain.AdapterBookingRepository {
	return &BookingRepositoryMysqlLayer{
		DB: db,
	}
}
