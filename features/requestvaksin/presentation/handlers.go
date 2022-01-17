package presentation

import (
	"SistemVaksinAPI/features/requestvaksin"
	requestvaksin_request "SistemVaksinAPI/features/requestvaksin/presentation/request"
	requestvaksin_response "SistemVaksinAPI/features/requestvaksin/presentation/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RequestvaksinHandler struct {
	requestvaksinBussiness requestvaksin.Bussiness
}

func NewRequestvaksinHandler(rbu requestvaksin.Bussiness) *RequestvaksinHandler {
	return &RequestvaksinHandler{
		requestvaksinBussiness: rbu,
	}
}

func (rh *RequestvaksinHandler) CreateRequestvaksin(c echo.Context) error {

	newRequestvaksin := requestvaksin_request.Requestvaksin{}

	c.Bind(&newRequestvaksin)

	result, err := rh.requestvaksinBussiness.CreateRequestvaksin(requestvaksin_request.ToCore(&newRequestvaksin))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "requestvaksin available",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    requestvaksin_response.FromCore(result),
	})
}

func (rh *RequestvaksinHandler) GetAllRequestvaksin(c echo.Context) error {
	result := rh.requestvaksinBussiness.GetAllRequestvaksin()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    requestvaksin_response.FromCoreSlice(result),
	})
}

func (rh *RequestvaksinHandler) GetRequestvaksinByID(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("ID"))

	result, _ := rh.requestvaksinBussiness.GetRequestvaksinByID(ID)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    requestvaksin_response.FromCore(result),
	})
}
