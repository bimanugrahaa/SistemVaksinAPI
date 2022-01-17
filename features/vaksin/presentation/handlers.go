package presentation

import (
	"SistemVaksinAPI/features/vaksin"
	vaksin_request "SistemVaksinAPI/features/vaksin/presentation/request"
	vaksin_response "SistemVaksinAPI/features/vaksin/presentation/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VaksinHandler struct {
	vaksinBussiness vaksin.Bussiness
}

func NewVaksinHandler(vbu vaksin.Bussiness) *VaksinHandler {
	return &VaksinHandler{
		vaksinBussiness: vbu,
	}
}

func (vh *VaksinHandler) CreateVaksin(c echo.Context) error {

	newVaksin := vaksin_request.Vaksin{}

	c.Bind(&newVaksin)

	result, err := vh.vaksinBussiness.CreateVaksin(vaksin_request.ToCore(newVaksin))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "vaksin available",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    vaksin_response.FromCore(result),
	})
}

func (vh *VaksinHandler) GetVaksinByID(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("ID"))

	result, _ := vh.vaksinBussiness.GetVaksinByID(ID)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    vaksin_response.FromCore(result),
	})
}

func (vh *VaksinHandler) UpdateVaksinByID(c echo.Context) error {
	updateVaksinByID := vaksin_request.Vaksin{}

	c.Bind(&updateVaksinByID)

	result, err := vh.vaksinBussiness.UpdateVaksinByID(vaksin_request.ToCore(updateVaksinByID))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    vaksin_response.FromCore(result),
	})
}
