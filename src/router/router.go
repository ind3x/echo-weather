package router

import (
	"echo-weather/src/api"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	citiesGroup := e.Group("/cities")

	//set main routes
	api.MainGroup(e)

	//set groupRoutes
	api.CitiesGroup(citiesGroup)

	return e
}