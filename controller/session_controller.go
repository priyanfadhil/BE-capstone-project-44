package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type EchoControllerSession struct {
	svc domain.AdapterSession
}

func (ce *EchoControllerSession) CreateSessionController(c echo.Context) error {
	session := model.Session{}
	c.Bind(&session)

	err := ce.svc.CreateSession(session)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"sessions": session,
	})
}

func (ce *EchoControllerSession) UpdateSessionController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	session := model.Session{}
	c.Bind(&session)

	err := ce.svc.UpdateSession(intID, session)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerSession) DeleteSessionController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteSessionByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerSession) GetSessionController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.svc.GetSessionByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"sessions": res,
	})
}

func (ce *EchoControllerSession) GetSessionsController(c echo.Context) error {
	session := ce.svc.GetAllSessions()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"sessions": session,
	}, "  ")
}
