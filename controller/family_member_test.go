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

func TestCreateFamilyMemberController(t *testing.T) {
	svc := MockFamilyMemberSvc{}

	svc.Mock.On("CreateFamilyMember", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("CreateFamilyMember", mock.Anything).
		Return(nil).Once()

	familyController := FamilyMemberController{
		svc: &svc,
	}

	t.Run("error", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		familyController.CreateFamilyMemberController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		familyController.CreateFamilyMemberController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetFamilyMemberyIDController(t *testing.T) {
	svc := MockFamilyMemberSvc{}

	svc.Mock.On("GetFamilyMemberyID", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetFamilyMemberyID", mock.Anything).
		Return(nil).Once()

	familyController := FamilyMemberController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		familyController.GetFamilyMemberController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetAllFamilyMemberController(t *testing.T) {
	svc := MockFamilyMemberSvc{}

	svc.Mock.On("GetAllFamilyMember", mock.Anything).
		Return(errors.New("new")).Once()

	svc.Mock.On("GetAllFamilyMember", mock.Anything).
		Return(nil).Once()

	familyController := FamilyMemberController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		familyController.GetFamilyMemberController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateFamilyMemberontroller(t *testing.T) {
	svc := MockFamilyMemberSvc{}

	svc.Mock.On("UpdateFamilyMember", mock.Anything).
		Return(nil).Once()

	familyController := FamilyMemberController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		familyController.UpdateFamilyMemberController(echoContext)

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteFamilyMemberontroller(t *testing.T) {
	svc := MockFamilyMemberSvc{}

	svc.Mock.On("DeleteFamilyMember", mock.Anything).
		Return(nil).Once()

	familyController := FamilyMemberController{
		svc: &svc,
	}

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		familyController.DeleteFamilyMemberController(echoContext)

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
