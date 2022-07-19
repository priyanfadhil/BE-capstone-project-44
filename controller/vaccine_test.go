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

func TestCreateVaccineController(t *testing.T) {
	svc := MockVaccineSvc{}

	svc.Mock.On("CreateVaccine", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateVaccine", mock.Anything).
		Return(nil).Once()

	vaccineController := VaccineController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccineController.CreateVaccineController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccineController.CreateVaccineController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetVaccineIDController(t *testing.T) {
	svc := MockVaccineSvc{}

	svc.Mock.On("GetVaccineyID", 1).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetVaccineyID", 1).
		Return(nil).Once()

	vaccineController := VaccineController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccineController.GetVaccineController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllVaccineController(t *testing.T) {
	svc := MockVaccineSvc{}

	svc.Mock.On("GetAllVaccine", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllVaccine", mock.Anything).
		Return(nil).Once()

	vaccineController := VaccineController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccineController.GetVaccineController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateVaccineontroller(t *testing.T) {
	svc := MockVaccineSvc{}

	svc.Mock.On("UpdateVaccine", mock.Anything).
		Return(nil).Once()

	vaccineController := VaccineController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccineController.UpdateVaccineController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteVaccineController(t *testing.T) {
	svc := MockVaccineSvc{}

	svc.Mock.On("DeleteVaccine", mock.Anything).
		Return(nil).Once()

	vaccineController := VaccineController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		vaccineController.DeleteVaccineController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
