package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/database"
	"github.com/priyanfadhil/BE-capstone-project-44/repository"
	"github.com/priyanfadhil/BE-capstone-project-44/usecase"
)

func VaccineStatusGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.VaccineStatusMysqlRepository(db)

	svcVaccineStatus := usecase.NewVaccineStatus(repo, conf)

	cont := EchoController{
		svc: svcVaccineStatus,
	}

	apiVaccineStatus := e.Group("/vaccinestatus",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiVaccineStatus.GET("", cont.GetVaccineStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccineStatus.GET("/:id", cont.GetVaccineStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccineStatus.PUT("/:id", cont.UpdateVaccineStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccineStatus.DELETE("/:id", cont.DeleteVaccineStatusController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccineStatus.POST("", cont.CreateVaccineStatusController)
}
