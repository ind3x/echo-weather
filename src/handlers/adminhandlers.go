package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCities(c echo.Context) error {
	return c.String(http.StatusOK, " You are on the Admin Page !!!")
}