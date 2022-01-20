package routes

import (
	"SistemVaksinAPI/config"
	"SistemVaksinAPI/factory"
	log "SistemVaksinAPI/middleware"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	presenter := factory.Init()
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	log.LogMiddleware(e)
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		SigningKey:    []byte(config.JwtSecret),
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"status":  "error login",
				"message": e,
			})
		},
	}))

	//Admins
	e.POST("/admin/register", presenter.AdminPresentation.CreateAdmin)
	e.POST("/admin/login", presenter.AdminPresentation.Login)

	//User
	e.POST("/user/register", presenter.UserPresentation.CreateUser)
	e.POST("/user/login", presenter.UserPresentation.Login)
	auth.PUT("/user", presenter.UserPresentation.UpdateUser)
	auth.GET("/user/:ID", presenter.UserPresentation.GetUserByID)

	//Faskes
	auth.POST("/faskes", presenter.FaskesPresentation.CreateFaskes)
	e.GET("/faskes/semua", presenter.FaskesPresentation.GetAllFaskes)
	e.GET("/faskes/:ID", presenter.FaskesPresentation.GetFaskesByID)
	e.GET("/faskesnama", presenter.FaskesPresentation.GetFaskesByName)

	//Vaksin
	auth.POST("/vaksin", presenter.VaksinPresentation.CreateVaksin)
	e.GET("/vaksin/:ID", presenter.VaksinPresentation.GetVaksinByID)
	auth.PUT("/vaksin", presenter.VaksinPresentation.UpdateVaksinByID)

	//Requestvaksin
	auth.POST("/requestvaksin", presenter.RequestvaksinPresentation.CreateRequestvaksin)
	e.GET("/requestvaksin", presenter.RequestvaksinPresentation.GetAllRequestvaksin)
	e.GET("/requestvaksin/:ID", presenter.RequestvaksinPresentation.GetRequestvaksinByID)
	e.GET("/requestvaksin/login", presenter.RequestvaksinPresentation.Login)
	auth.PUT("/status/satu", presenter.RequestvaksinPresentation.UpdateRequestVaksinSatu)
	auth.PUT("/status/dua", presenter.RequestvaksinPresentation.UpdateRequestVaksinDua)

	return e
}
