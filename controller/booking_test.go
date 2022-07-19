package controller

import (
	// "errors"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/mock"
)

func TestCreateBookingController(t *testing.T) {
	svc := MockBookingSvc{}

	svc.Mock.On("CreateBooking", 0, mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateBooking", 0, mock.Anything).
		Return(nil).Once()

	usrController := BookingController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateBookingController(echoContext)

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateBookingController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetBookingByIDController(t *testing.T) {
	svc := MockBookingSvc{}

	svc.Mock.On("GetBookingByID", 1).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetBookingByID", 1).
		Return(nil).Once()

	usrController := BookingController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetBookingController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllBookingsController(t *testing.T) {
	svc := MockBookingSvc{}

	svc.Mock.On("GetAllBookings", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllBookings", mock.Anything).
		Return(nil).Once()

	usrController := BookingController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetBookingsController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateBookingController(t *testing.T) {
	svc := MockBookingSvc{}

	svc.Mock.On("UpdateBooking", 0, mock.Anything).
		Return(nil).Once()

	usrController := BookingController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.UpdateBookingController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteBookingController(t *testing.T) {
	svc := MockBookingSvc{}

	svc.Mock.On("DeleteBooking", mock.Anything).
		Return(nil).Once()

	usrController := BookingController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteBookingController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
