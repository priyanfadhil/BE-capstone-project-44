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

func TestCreateUserController(t *testing.T) {
	svc := MockUserSvc{}

	svc.Mock.On("CreateUser", 0, mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateUser", 0, mock.Anything).
		Return(nil).Once()

	usrController := UserController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetUserIDController(t *testing.T) {
	svc := MockUserSvc{}

	svc.Mock.On("GetUseryID", 1).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetUseryID", 1).
		Return(nil).Once()

	usrController := UserController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetUserController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllUserController(t *testing.T) {
	svc := MockUserSvc{}

	svc.Mock.On("GetAllUser", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllUser", mock.Anything).
		Return(nil).Once()

	usrController := UserController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.GetUserController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateUserontroller(t *testing.T) {
	svc := MockUserSvc{}

	svc.Mock.On("UpdateUser", 0, mock.Anything).
		Return(nil).Once()

	usrController := UserController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.UpdateUserController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteUserController(t *testing.T) {
	svc := MockUserSvc{}

	svc.Mock.On("DeleteUser", mock.Anything).
		Return(nil).Once()

	usrController := UserController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteUserController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
