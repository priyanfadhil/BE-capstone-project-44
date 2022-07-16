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

func TestCreateSessionController(t *testing.T) {
	svc := MockSessionSvc{}

	svc.Mock.On("CreateSession", 1, mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateSession", 1, mock.Anything).
		Return(nil).Once()

	sesiController := SessionController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		sesiController.CreateSessionController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		sesiController.CreateSessionController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetSessionIDController(t *testing.T) {
	svc := MockSessionSvc{}

	svc.Mock.On("GetSessionyID", 1).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetSessionyID", 1).
		Return(nil).Once()

	sesiController := SessionController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		sesiController.GetSessionController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllSessionController(t *testing.T) {
	svc := MockSessionSvc{}

	svc.Mock.On("GetAllSession", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllSession", mock.Anything).
		Return(nil).Once()

	sesiController := SessionController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		sesiController.GetSessionController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateSessionontroller(t *testing.T) {
	svc := MockSessionSvc{}

	svc.Mock.On("UpdateSession", 0, mock.Anything).
		Return(nil).Once()

	sesiController := SessionController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		sesiController.UpdateSessionController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteSessionController(t *testing.T) {
	svc := MockSessionSvc{}

	svc.Mock.On("DeleteSession", mock.Anything).
		Return(nil).Once()

	sesiController := SessionController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		sesiController.DeleteSessionController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
