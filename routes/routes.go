package routes

import (
	"SistemVaksinAPI/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	presenter := factory.Init()
	e := echo.New()
	e.Use(middleware.CORS())

	//Admins
	e.POST("/admin/register", presenter.AdminPresentation.CreateAdmin)
	e.POST("/admin/login", presenter.AdminPresentation.Login)

	//User
	e.POST("/user/register", presenter.UserPresentation.CreateUser)
	e.POST("/user/login", presenter.UserPresentation.Login)
	e.PUT("/user", presenter.UserPresentation.UpdateUser)
	e.GET("/user", presenter.UserPresentation.GetUserByID)

	//Faskes
	e.POST("/faskes", presenter.FaskesPresentation.CreateFaskes)
	e.GET("/faskes/semua", presenter.FaskesPresentation.GetAllFaskes)
	e.GET("/faskes/:ID", presenter.FaskesPresentation.GetFaskesByID)
	e.GET("/faskesnama", presenter.FaskesPresentation.GetFaskesByName)

	//Vaksin
	e.POST("/vaksin", presenter.VaksinPresentation.CreateVaksin)
	e.GET("/vaksin/:ID", presenter.VaksinPresentation.GetVaksinByID)
	e.PUT("/vaksin", presenter.VaksinPresentation.UpdateVaksinByID)

	//Requestvaksin
	e.POST("/requestvaksin", presenter.RequestvaksinPresentation.CreateRequestvaksin)
	e.GET("/requestvaksin", presenter.RequestvaksinPresentation.GetAllRequestvaksin)
	e.GET("/requestvaksin/:ID", presenter.RequestvaksinPresentation.GetRequestvaksinByID)
	e.GET("/requestvaksin/login", presenter.RequestvaksinPresentation.Login)

	e.POST("/request", presenter.RequestvaksinPresentation.CreateRequestvaksin)
	e.GET("/request", presenter.RequestvaksinPresentation.GetAllRequestvaksin)
	e.GET("/request/:ID", presenter.RequestvaksinPresentation.GetRequestvaksinByID)
	return e
}
