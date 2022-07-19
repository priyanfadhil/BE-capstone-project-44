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

func TestCreateVaccinationLocationController(t *testing.T) {
	svc := MockVaccinationLocationSvc{}

	svc.Mock.On("CreateVaccinationLocation", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateVaccinationLocation", mock.Anything).
		Return(nil).Once()

	admController := VaccinationLocationController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.CreateVaccinationLocationController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.CreateVaccinationLocationController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetVaccinationLocationIDController(t *testing.T) {
	svc := MockVaccinationLocationSvc{}

	svc.Mock.On("GetVaccinationLocationyID", 1).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetVaccinationLocationyID", 1).
		Return(nil).Once()

	admController := VaccinationLocationController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.GetVaccinationLocationController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllVaccinationLocationController(t *testing.T) {
	svc := MockVaccinationLocationSvc{}

	svc.Mock.On("GetAllVaccinationLocation", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllVaccinationLocation", mock.Anything).
		Return(nil).Once()

	admController := VaccinationLocationController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.GetVaccinationLocationController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateVaccinationLocationontroller(t *testing.T) {
	svc := MockVaccinationLocationSvc{}

	svc.Mock.On("UpdateVaccinationLocation", mock.Anything).
		Return(nil).Once()

	admController := VaccinationLocationController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.UpdateVaccinationLocationController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteVaccinationLocationController(t *testing.T) {
	svc := MockVaccinationLocationSvc{}

	svc.Mock.On("DeleteVaccinationLocation", mock.Anything).
		Return(nil).Once()

	admController := VaccinationLocationController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.DeleteVaccinationLocationController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
