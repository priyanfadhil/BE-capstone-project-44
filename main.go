package main

import (
	"github.com/labstack/echo/v4"
	c "github.com/priyanfadhil/BE-capstone-project-44/config"
	rest "github.com/priyanfadhil/BE-capstone-project-44/controller"
)

func main() {
	config := c.InitConfiguration()
	e := echo.New()

	rest.UserGroupAPI(e, config)
	rest.FamilyMemberGroupAPI(e, config)
	rest.BookingGroupAPI(e, config)
	rest.SessionGroupAPI(e, config)
	rest.VaccineStatusGroupAPI(e, config)
	rest.UserAddressGroupAPI(e, config)
	rest.VaccinationLocationGroupAPI(e, config)
	rest.VaccineGroupAPI(e, config)

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
