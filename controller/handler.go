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

	Userrepo := repository.UserMysqlRepository(db)
	Familymember := repository.FamilyMemberMysqlRepository(db)

	svcUser := usecase.User(Userrepo, Familymember, conf)

	cont := UserController{
		svc: svcUser,
	}

	e.GET("/debug", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "hello world",
		})
	})

	apiUser := e.Group("/users",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiUser.GET("", cont.GetUsersController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("/login", cont.LoginUserController)
	apiUser.GET("/:id", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/:id", cont.UpdateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.DELETE("/:id", cont.DeleteUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("", cont.CreateUserController)
}

func FamilyMemberGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.FamilyMemberMysqlRepository(db)

	svcFamilyMember := usecase.FamilyMember(repo, conf)

	cont := FamilyMemberController{
		svc: svcFamilyMember,
	}

	apiFamilyMember := e.Group("/familymembers",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiFamilyMember.GET("", cont.GetFamilyMembersController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiFamilyMember.GET("/:id", cont.GetFamilyMemberController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiFamilyMember.PUT("/:id", cont.UpdateFamilyMemberController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiFamilyMember.DELETE("/:id", cont.DeleteFamilyMemberController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiFamilyMember.POST("", cont.CreateFamilyMemberController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func BookingGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	bookingRepo := repository.NewBookingMysqlRepository(db)
	sessionRepo := repository.NewSessionMysqlRepository(db)

	svcBooking := usecase.NewBooking(bookingRepo, sessionRepo, conf)

	cont := BookingController{
		svc: svcBooking,
	}

	apiBooking := e.Group("/bookings",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiBooking.GET("", cont.GetBookingsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiBooking.GET("/:id", cont.GetBookingController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiBooking.PUT("/:id", cont.UpdateBookingController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiBooking.DELETE("/:id", cont.DeleteBookingController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiBooking.POST("", cont.CreateBookingController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func SessionGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.NewSessionMysqlRepository(db)

	svcSession := usecase.NewSession(repo, conf)

	cont := EchoControllerSession{
		svc: svcSession,
	}

	apiSession := e.Group("/sessions",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiSession.GET("", cont.GetSessionsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiSession.GET("/:id", cont.GetSessionController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiSession.PUT("/:id", cont.UpdateSessionController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiSession.DELETE("/:id", cont.DeleteSessionController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiSession.POST("", cont.CreateSessionController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func VaccinationLocationGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.VaccinationLocationMysqlRepository(db)

	svcVaccinationLocation := usecase.NewVaccinationLocation(repo, conf)

	cont := VaccinationLocationController{
		svc: svcVaccinationLocation,
	}

	apiVaccinationLocation := e.Group("/vaccinationlocations",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiVaccinationLocation.GET("", cont.GetVaccinationLocationsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccinationLocation.GET("/:id", cont.GetVaccinationLocationController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccinationLocation.PUT("/:id", cont.UpdateVaccinationLocationController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccinationLocation.DELETE("/:id", cont.DeleteVaccinationLocationController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccinationLocation.POST("", cont.CreateVaccinationLocationController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func VaccineGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.NewVaccineMysqlRepository(db)

	svcVaccine := usecase.NewVaccine(repo, conf)

	cont := VaccineController{
		svc: svcVaccine,
	}

	apiVaccine := e.Group("/vaccines",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiVaccine.GET("", cont.GetAllVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccine.GET("/:id", cont.GetVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccine.PUT("/:id", cont.UpdateVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccine.DELETE("/:id", cont.DeleteVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiVaccine.POST("", cont.CreateVaccineController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func AdminGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.NewAdminMysqlRepository(db)

	svcAdmin := usecase.NewAdmin(repo, conf)

	cont := AdminController{
		svc: svcAdmin,
	}

	apiAdmin := e.Group("/admins",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiAdmin.GET("", cont.GetAllAdminController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiAdmin.POST("/login", cont.LoginAdminController)
	apiAdmin.GET("/:id", cont.GetAdminController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiAdmin.PUT("/:id", cont.UpdateAdminController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiAdmin.DELETE("/:id", cont.DeleteAdminController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiAdmin.POST("", cont.CreateAdminController)
}
