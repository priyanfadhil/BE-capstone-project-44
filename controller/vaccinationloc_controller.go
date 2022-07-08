package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type EchoControllerVaccinationLocation struct {
	svc domain.AdapterVaccinationLocation

func (ce *EchoControllerVaccinationLocation) CreateVaccinationLocationController(c echo.Context) error {
	vaccinationlocation := model.VaccinationLocation{}
	c.Bind(&vaccinationlocation)

	err := ce.svc.VaccinationLocation(vaccinationlocation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages":       "success",
		"vaccinationlocation": vaccinationlocation,
	})
}

func (ce *EchoControllerVaccinationLocation) UpdateVaccinationLocationController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	vaccinationlocation := model.VaccinationLocation{}
	c.Bind(&vaccinationlocation)

	err := ce.svc.UpdateVaccinationLocation(intID, vaccinationlocation)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerVaccinationLocation) DeleteVaccinationLocationController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteVaccinationLocationByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerVaccinationLocation) GetVaccinationLocationController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.svc.GetVaccinationLocationByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":       "success",
		"vaccinationlocation": res,
	})
}

func (ce *EchoControllerVaccinationLocation) GetVaccinationLocationController(c echo.Context) error {
	vaccinationlocation := ce.svc.GetAllVaccinationLocation()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":       "success",
		"vaccinationlocation": vaccinationlocation,
	}, "  ")
}
