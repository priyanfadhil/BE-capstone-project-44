package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/database"
	"github.com/priyanfadhil/BE-capstone-project-44/repository"
	"github.com/priyanfadhil/BE-capstone-project-44/usecase"
)

func VaccineGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.VaccineMysqlRepository(db)

	svcVaccine := usecase.NewVaccine(repo, conf)

	cont := EchoController{
		svc: svcVaccine,
	}

	apiVaccine := e.Group("/vaccine",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiVaccine.GET("", cont.GetVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccine.GET("/:id", cont.GetVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccine.PUT("/:id", cont.UpdateVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccine.DELETE("/:id", cont.DeleteVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccine.POST("", cont.CreateVaccineController)
}
