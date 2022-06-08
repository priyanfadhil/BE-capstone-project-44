package main

import (
	"github.com/labstack/echo/v4"
	c "github.com/priyanfadhil/BE-capstone-project-44/config"
)

func main() {
	config := c.InitConfiguration()
	e := echo.New()

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
