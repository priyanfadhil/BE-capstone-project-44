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

func TestCreateAdminController(t *testing.T) {
	svc := MockAdminSvc{}

	svc.Mock.On("CreateAdmin", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateAdmin", mock.Anything).
		Return(nil).Once()

	admController := AdminController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.CreateAdminController(echoContext)

		assert.Equal(t, 400, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.CreateAdminController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetAdminIDController(t *testing.T) {
	svc := MockAdminSvc{}

	svc.Mock.On("GetAdminyID", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAdminyID", mock.Anything).
		Return(nil).Once()

	admController := AdminController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.GetAdminController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllAdminController(t *testing.T) {
	svc := MockAdminSvc{}

	svc.Mock.On("GetAllAdmin", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllAdmin", mock.Anything).
		Return(nil).Once()

	admController := AdminController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.GetAdminController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateAdminontroller(t *testing.T) {
	svc := MockAdminSvc{}

	svc.Mock.On("UpdateAdmin", mock.Anything).
		Return(nil).Once()

	admController := AdminController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.UpdateAdminController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteAdminController(t *testing.T) {
	svc := MockAdminSvc{}

	svc.Mock.On("DeleteAdmin", mock.Anything).
		Return(nil).Once()

	admController := AdminController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		admController.DeleteAdminController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
