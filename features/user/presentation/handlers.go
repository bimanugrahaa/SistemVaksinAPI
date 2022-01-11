package presentation

import (
	"SistemVaksinAPI/features/user"
	"net/http"

	user_request "SistemVaksinAPI/features/user/presentation/request"
	user_response "SistemVaksinAPI/features/user/presentation/response"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBussiness user.Bussiness
}

func NewUserHandler(ubu user.Bussiness) *UserHandler {
	return &UserHandler{
		userBussiness: ubu,
	}
}

func (uh *UserHandler) CreateUser(c echo.Context) error {

	newUser := user_request.User{}

	c.Bind(&newUser)

	result, err := uh.userBussiness.CreateUser(user_request.ToCore(newUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username available",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    user_response.FromCore(result),
	})
}

func (uh *UserHandler) Login(c echo.Context) error {

	infoUser := user_request.User{}

	c.Bind(&infoUser)

	result, err := uh.userBussiness.Login(user_request.ToCore(infoUser))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "something went wrong",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    user_response.FromCoreLogin(result),
	})
}
