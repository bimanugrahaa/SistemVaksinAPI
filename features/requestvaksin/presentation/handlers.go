package presentation

import (
	"SistemVaksinAPI/features/requestvaksin"
	requestvaksin_request "SistemVaksinAPI/features/requestvaksin/presentation/request"
	requestvaksin_response "SistemVaksinAPI/features/requestvaksin/presentation/response"
	"fmt"
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

	fmt.Println("newRequestvaksin", newRequestvaksin)
	result, err := rh.requestvaksinBussiness.CreateRequestvaksin(requestvaksin_request.ToCore(newRequestvaksin))
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

	result, err := rh.requestvaksinBussiness.GetRequestvaksinByID(ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "requestvaksin not available",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    requestvaksin_response.FromCore(result),
	})
}

func (rh *RequestvaksinHandler) GetRequestvaksinByVaksinID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("ID"))
	// infoRequestvaksin := requestvaksin_request.Requestvaksin{}

	// c.Bind(&infoRequestvaksin)
	// id := infoRequestvaksin.ID
	// fmt.Println("id", id)
	result, err := rh.requestvaksinBussiness.GetRequestvaksinByVaksinID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "requestvaksin not available",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    result,
	})
}

func (rh *RequestvaksinHandler) Login(c echo.Context) error {

	infoRequestvaksin := requestvaksin_request.Requestvaksin{}

	c.Bind(&infoRequestvaksin)

	result, err := rh.requestvaksinBussiness.Login(requestvaksin_request.ToCore(infoRequestvaksin))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "something went wrong",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    requestvaksin_response.FromCore(result),
	})
}

func (rh *RequestvaksinHandler) UpdateRequestVaksinSatu(c echo.Context) error {
	UpdateRequestVaksinSatu := requestvaksin_request.Requestvaksin{}

	c.Bind(&UpdateRequestVaksinSatu)

	result, err := rh.requestvaksinBussiness.UpdateRequestVaksinSatu(requestvaksin_request.ToCore(UpdateRequestVaksinSatu))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    requestvaksin_response.FromCore(result),
	})
}

func (rh *RequestvaksinHandler) UpdateRequestVaksinDua(c echo.Context) error {
	UpdateRequestVaksinSatu := requestvaksin_request.Requestvaksin{}

	c.Bind(&UpdateRequestVaksinSatu)

	result, err := rh.requestvaksinBussiness.UpdateRequestVaksinDua(requestvaksin_request.ToCore(UpdateRequestVaksinSatu))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    requestvaksin_response.FromCore(result),
	})
}
