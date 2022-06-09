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

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
