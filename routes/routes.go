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

	//Faskes
	e.POST("/faskes", presenter.FaskesPresentation.CreateFaskes)
	e.GET("/faskes/semua", presenter.FaskesPresentation.GetAllFaskes)
	e.GET("/faskes/:ID", presenter.FaskesPresentation.GetFaskesByID)

	//Vaksin
	e.POST("/vaksin", presenter.VaksinPresentation.CreateVaksin)
	e.GET("/vaksin/:ID", presenter.VaksinPresentation.GetVaksinByID)
	e.PUT("/vaksin", presenter.VaksinPresentation.UpdateVaksinByID)

	return e
}
