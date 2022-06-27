package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type UserAddressController struct {
	svc domain.AdapterUserAddress
}

func (ce *UserAddressController) CreateUserAddressController(c echo.Context) error {
	useraddress := model.UserAddress{}
	c.Bind(&useraddress)

	err := ce.svc.CreateUserAddress(useraddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages":     "success",
		"useraddresss": useraddress,
	})
}

func (ce *UserAddressController) UpdateUserAddressController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	useraddress := model.UserAddress{}
	c.Bind(&useraddress)

	err := ce.svc.UpdateUserAddress(intID, useraddress)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *UserAddressController) DeleteUserAddressController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteUserAddressByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *UserAddressController) GetUserAddressController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.svc.GetUserAddressByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":     "success",
		"useraddresss": res,
	})
}

func (ce *UserAddressController) GetAllUserAddressController(c echo.Context) error {
	useraddress := ce.svc.GetAllUserAddress()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":     "success",
		"useraddresss": useraddress,
	}, "  ")
}
