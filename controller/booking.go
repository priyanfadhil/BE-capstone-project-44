package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type BookingController struct {
	svc domain.AdapterBooking
}

func (ce *BookingController) CreateBookingController(c echo.Context) error {
	booking := model.Booking{}
	c.Bind(&booking)
	session_id := booking.SessionID

	err := ce.svc.CreateBooking(session_id, booking)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"bookings": booking,
	})
}

func (ce *BookingController) UpdateBookingController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	booking := model.Booking{}
	c.Bind(&booking)

	err := ce.svc.UpdateBooking(intID, booking)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *BookingController) DeleteBookingController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteBookingByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *BookingController) GetBookingController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.svc.GetBookingByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"bookings": res,
	})
}

func (ce *BookingController) GetBookingsController(c echo.Context) error {
	booking := ce.svc.GetAllBookings()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"bookings": booking,
	}, "  ")
}
