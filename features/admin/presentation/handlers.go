package presentation

import (
	"SistemVaksinAPI/features/admin"
	admin_request "SistemVaksinAPI/features/admin/presentation/request"
	admin_response "SistemVaksinAPI/features/admin/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminBussiness admin.Bussiness
}

func NewAdminHandler(abu admin.Bussiness) *AdminHandler {
	return &AdminHandler{
		adminBussiness: abu,
	}
}

func (ah *AdminHandler) CreateAdmin(c echo.Context) error {

	newAdmin := admin_request.Admin{}

	c.Bind(&newAdmin)

	result, err := ah.adminBussiness.CreateAdmin(admin_request.ToCore(newAdmin))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username available",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    admin_response.FromCore(result),
	})
}
