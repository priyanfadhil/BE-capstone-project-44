package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type VaccineController struct {
	svc domain.AdapterVaccine
}

func (ce *VaccineController) CreateVaccineController(c echo.Context) error {
	vaccine := model.Vaccine{}
	c.Bind(&vaccine)

	err := ce.svc.CreateVaccine(vaccine)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"vaccine":  vaccine,
	})
}

func (ce *VaccineController) UpdateVaccineController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	vaccine := model.Vaccine{}
	c.Bind(&vaccine)

	err := ce.svc.UpdateVaccine(intID, vaccine)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *VaccineController) DeleteVaccineController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteVaccineByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *VaccineController) GetVaccineController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.svc.GetVaccineByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"vaccine":  res,
	})
}

func (ce *VaccineController) GetAllVaccineController(c echo.Context) error {
	vaccine := ce.svc.GetAllVaccines()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"vaccine":  vaccine,
	}, "  ")
}
