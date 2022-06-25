package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type VaccinationLocationController struct {
	svc domain.AdapterVaccinationLocation
}

func (ce *VaccinationLocationController) CreateVaccinationLocationController(c echo.Context) error {
	vaccinationlocation := model.VaccinationLocation{}
	c.Bind(&vaccinationlocation)

	err := ce.svc.CreateVaccinationLocation(vaccinationlocation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages":             "success",
		"vaccinationlocations": vaccinationlocation,
	})
}

func (ce *VaccinationLocationController) UpdateVaccinationLocationController(c echo.Context) error {
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

func (ce *VaccinationLocationController) DeleteVaccinationLocationController(c echo.Context) error {
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

func (ce *VaccinationLocationController) GetVaccinationLocationController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.svc.GetVaccinationLocationByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":             "success",
		"vaccinationlocations": res,
	})
}

func (ce *VaccinationLocationController) GetVaccinationLocationsController(c echo.Context) error {
	vaccinationlocations := ce.svc.GetAllVaccinationLocations()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":             "success",
		"vaccinationlocations": vaccinationlocations,
	}, "  ")
}
