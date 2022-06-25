package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type EchoControllerVaccineStatus struct {
	svc domain.AdapterVaccineStatus
}

func (ce *EchoControllerVaccineStatus) CreateVaccineStatusController(c echo.Context) error {
	vaccinestatus := model.VaccineStatus{}
	c.Bind(&vaccinestatus)

	err := ce.svc.VaccineStatus(vaccinestatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages":       "success",
		"vaccinesstatus": vaccinestatus,
	})
}

func (ce *EchoControllerVaccineStatus) UpdateVaccineStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	vaccinestatus := model.VaccineStatus{}
	c.Bind(&vaccinestatus)

	err := ce.svc.UpdateVaccineStatus(intID, vaccinestatus)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerVaccineStatus) DeleteVaccineStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteVaccineStatusByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerVaccineStatus) GetVaccineStatusController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {

	}

	res, err := ce.svc.GetVaccineStatusByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":       "success",
		"vaccinesstatus": res,
	})
}

func (ce *EchoControllerVaccineStatus) GetVaccineStatusController(c echo.Context) error {
	vaccinestatus := ce.svc.GetAllVaccineStatus()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":       "success",
		"vaccinesstatus": vaccinestatus,
	}, "  ")
}
