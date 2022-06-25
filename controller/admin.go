package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type AdminController struct {
	svc domain.AdapterAdmin
}

func (ce *AdminController) CreateAdminController(c echo.Context) error {
	admin := model.Admin{}
	c.Bind(&admin)

	err := ce.svc.CreateAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"admin":    admin,
	})
}

func (ce *AdminController) UpdateAdminController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	admin := model.Admin{}
	c.Bind(&admin)

	err := ce.svc.UpdateAdmin(intID, admin)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *AdminController) DeleteAdminController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteAdminByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *AdminController) GetAdminController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.svc.GetAdminByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"admin":    res,
	})
}

func (ce *AdminController) GetAllAdminController(c echo.Context) error {
	admin := ce.svc.GetAllAdmins()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"admin":    admin,
	}, "  ")
}
