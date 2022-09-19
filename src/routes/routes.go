package routes

import (
	"github.com/labstack/echo/v4"
	"softwareLab/src/api"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/register", api.RegisterView)
	e.POST("/register", api.RegisterView)
	e.POST("/registerUser", api.RegisterUser)
	e.GET("/home", api.Home)
	e.POST("/home", api.Home)
	e.POST("/login", api.Login)
	e.GET("/mentee", api.MenteeDashboard)
	e.POST("/mentee", api.MenteeDashboard)
	e.GET("/mentee/account", api.MenteeAccount)
	e.POST("/mentee/account/process", api.MenteeAccountProcess)
	e.GET("/mentee/profile", api.MenteeProfile)
	e.POST("/mentee/profile", api.MenteeProfile)
	e.GET("/mentor/register", api.MentorRegister)
	e.POST("/mentor/register", api.MentorRegister)
	e.POST("/register/mentor/process", api.MentorRegisterProcess)
	e.GET("/mentor/dashboard", api.MentorDashboard)
	e.POST("/mentor/dashboard", api.MentorDashboard)
	e.GET("/mentor/login", api.MentorLogin)
	e.POST("/mentor/login/process", api.MentorLoginProcess)
	e.GET("/room", api.Room)
	e.POST("/room", api.Room)
}
