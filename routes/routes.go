package routes

import (
	"SistemVaksinAPI/factory"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	presenter := factory.Init()
	e := echo.New()

	//Admins
	e.POST("/admin/register", presenter.AdminPresentation.CreateAdmin)
	e.POST("/admin/login", presenter.AdminPresentation.Login)

	e.POST("/user/register", presenter.UserPresentation.CreateUser)
	e.POST("/user/login", presenter.UserPresentation.Login)

	return e
}
