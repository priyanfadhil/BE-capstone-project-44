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

func TestCreateVaccineStatusController(t *testing.T) {
	svc := MockVaccineStatusSvc{}

	svc.Mock.On("CreateVaccineStatus", 0, mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateVaccineStatus", 0, mock.Anything).
		Return(nil).Once()

	vaccinestatusController := VaccineStatusController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccinestatusController.CreateVaccineStatusController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccinestatusController.CreateVaccineStatusController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetVaccineStatusByIDController(t *testing.T) {
	svc := MockVaccineStatusSvc{}

	svc.Mock.On("GetVaccineStatusByID", 1).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetVaccineStatusByID", 1).
		Return(nil).Once()

	vaccinestatusController := VaccineStatusController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccinestatusController.GetVaccineStatusController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllVaccineStatusController(t *testing.T) {
	svc := MockVaccineStatusSvc{}

	svc.Mock.On("GetAllVaccineStatus", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllVaccineStatus", mock.Anything).
		Return(nil).Once()

	vaccinestatusController := VaccineStatusController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccinestatusController.GetAllVaccineStatusesController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateVaccineStatusController(t *testing.T) {
	svc := MockVaccineStatusSvc{}

	svc.Mock.On("UpdateVaccineStatus", 0, mock.Anything).
		Return(nil).Once()

	vaccinestatusController := VaccineStatusController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccinestatusController.UpdateVaccineStatusController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteVaccineStatusController(t *testing.T) {
	svc := MockVaccineStatusSvc{}

	svc.Mock.On("DeleteVaccineStatus", mock.Anything).
		Return(nil).Once()

	vaccinestatusController := VaccineStatusController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccinestatusController.DeleteVaccineStatusController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
