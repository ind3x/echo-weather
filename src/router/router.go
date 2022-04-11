package router

import (
	"echo-weather/src/api"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//set main routes
	api.MainGroup(e)

	return e
}