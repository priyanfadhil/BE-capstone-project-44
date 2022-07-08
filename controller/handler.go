package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/priyanfadhil/BE-capstone-project-44/config"
	"github.com/priyanfadhil/BE-capstone-project-44/database"
	"github.com/priyanfadhil/BE-capstone-project-44/repository"
	"github.com/priyanfadhil/BE-capstone-project-44/usecase"
)

func UserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.UserMysqlRepository(db)

	svcUser := usecase.NewUser(repo, conf)

	cont := EchoController{
		svc: svcUser,
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	apiUser := e.Group("/users",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)
	apiUser.GET("", cont.GetUsersController)
	apiUser.GET("/:id", cont.GetUserController)
	apiUser.PUT("/:id", cont.UpdateUserController)
	apiUser.DELETE("/:id", cont.DeleteUserController)
	apiUser.POST("", cont.CreateUserController)
}

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

	apiVaccineStatus.GET("", cont.GetVaccineStatusController)
	apiVaccineStatus.GET("/:id", cont.GetVaccineStatusController)
	apiVaccineStatus.PUT("/:id", cont.UpdateVaccineStatusController)
	apiVaccineStatus.DELETE("/:id", cont.DeleteVaccineStatusController)
	apiVaccineStatus.POST("", cont.CreateVaccineStatusController)
}
