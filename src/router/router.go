package router

import (
	"echo-weather/src/api"
	"echo-weather/src/api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")

	//set all middlewares
	middlewares.SetMainMiddleWares(e)
	middlewares.SetAdminMiddlewares(adminGroup)

	//set main routes
	api.MainGroup(e)

	//set groupRoutes
	api.AdminGroup(adminGroup)
	return e
}