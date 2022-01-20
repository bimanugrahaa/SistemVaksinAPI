package presentation

import (
	"SistemVaksinAPI/features/faskes"
	"fmt"
	"net/http"
	"strconv"

	faskes_request "SistemVaksinAPI/features/faskes/presentation/request"
	faskes_response "SistemVaksinAPI/features/faskes/presentation/response"

	"github.com/labstack/echo/v4"
)

type FaskesHandler struct {
	faskesBussiness faskes.Bussiness
}

func NewFaskesHandler(fbu faskes.Bussiness) *FaskesHandler {
	return &FaskesHandler{
		faskesBussiness: fbu,
	}
}

func (fh *FaskesHandler) CreateFaskes(c echo.Context) error {

	newFaskes := faskes_request.Faskes{}

	c.Bind(&newFaskes)

	fmt.Println(newFaskes)

	result, err := fh.faskesBussiness.CreateFaskes(faskes_request.ToCore(newFaskes))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "faskes available",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    faskes_response.FromCore(result),
	})
}

func (fh *FaskesHandler) GetAllFaskes(c echo.Context) error {
	result := fh.faskesBussiness.GetAllFaskes()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    faskes_response.FromCoreSlice(result),
	})
}

func (fh *FaskesHandler) GetFaskesByID(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("ID"))

	result, _ := fh.faskesBussiness.GetFaskesByID(ID)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    faskes_response.FromCore(result),
	})
}
