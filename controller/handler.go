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

	svcUser := usecase.User(repo, conf)

	cont := UserController{
		svc: svcUser,
	}

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

	repo := repository.NewBookingMysqlRepository(db)

	svcBooking := usecase.NewBooking(repo, conf)

	cont := EchoControllerBooking{
		svc: svcBooking,
	}

	apiBooking := e.Group("/booking",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiBooking.GET("", cont.GetBookingsController)
	apiBooking.GET("/:id", cont.GetBookingController)
	apiBooking.PUT("/:id", cont.UpdateBookingController)
	apiBooking.DELETE("/:id", cont.DeleteBookingController)
	apiBooking.POST("", cont.CreateBookingController)
}

func SessionGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)

	repo := repository.NewSessionMysqlRepository(db)

	svcSession := usecase.NewSession(repo, conf)

	cont := EchoControllerSession{
		svc: svcSession,
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	apiSession := e.Group("/session",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiSession.GET("", cont.GetSessionsController)
	apiSession.GET("/:id", cont.GetSessionController)
	apiSession.PUT("/:id", cont.UpdateSessionController)
	apiSession.DELETE("/:id", cont.DeleteSessionController)
	apiSession.POST("", cont.CreateSessionController)
}
