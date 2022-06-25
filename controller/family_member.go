package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/priyanfadhil/BE-capstone-project-44/domain"
	"github.com/priyanfadhil/BE-capstone-project-44/model"
)

type FamilyMemberController struct {
	svc domain.AdapterFamilyMember
}

func (ce *FamilyMemberController) CreateFamilyMemberController(c echo.Context) error {
	familymember := model.FamilyMember{}
	c.Bind(&familymember)

	err := ce.svc.CreateFamilyMember(familymember)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages":      "success",
		"familymembers": familymember,
	})
}

func (ce *FamilyMemberController) UpdateFamilyMemberController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	familymember := model.FamilyMember{}
	c.Bind(&familymember)

	bearer := c.Get("familymember").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	err := ce.svc.UpdateFamilyMember(intID, int(claim["id"].(float64)), familymember)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":   "edited",
		"id":         intID,
		"expeted id": claim["id"],
	})
}

func (ce *FamilyMemberController) DeleteFamilyMemberController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteFamilyMemberByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *FamilyMemberController) GetFamilyMemberController(c echo.Context) error {
	fmt.Println("eksekusi handler")
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.svc.GetFamilyMemberByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":      "success",
		"familymembers": res,
	})
}

func (ce *FamilyMemberController) GetFamilyMembersController(c echo.Context) error {
	familymembers := ce.svc.GetAllFamilyMembers()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":      "success",
		"familymembers": familymembers,
	}, "  ")
}
