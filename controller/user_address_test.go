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

func TestCreateUserAddressController(t *testing.T) {
	svc := MockUserAddressSvc{}

	svc.Mock.On("CreateUserAddress", 0, mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateUserAddress", 0, mock.Anything).
		Return(nil).Once()

	usrController := UserAddressController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserAddressController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserAddressController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetUserAddressByIDController(t *testing.T) {
	svc := MockUserAddressSvc{}

	svc.Mock.On("GetUserAddressByID", 1).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetUserAddressByID", 1).
		Return(nil).Once()

	usrController := UserAddressController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetUserAddressController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllUserAddressesController(t *testing.T) {
	svc := MockUserAddressSvc{}

	svc.Mock.On("GetAllUserAddresses", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllUserAddresses", mock.Anything).
		Return(nil).Once()

	usrController := UserAddressController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetAllUserAddressesController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateUserAddressController(t *testing.T) {
	svc := MockUserAddressSvc{}

	svc.Mock.On("UpdateUserAddress", 0, mock.Anything).
		Return(nil).Once()

	usrController := UserAddressController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.UpdateUserAddressController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteUserAddressController(t *testing.T) {
	svc := MockUserAddressSvc{}

	svc.Mock.On("DeleteUserAddress", mock.Anything).
		Return(nil).Once()

	usrController := UserAddressController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteUserAddressController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
