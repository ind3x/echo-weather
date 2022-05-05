package api

import (
	"echo-weather/src/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/healthy", handlers.HealthCheck)
	// e.GET("/cities", handlers.GetCities)
}

func CitiesGroup(g *echo.Group) {
	g.GET("", handlers.GetCities)
	g.GET("/:q", handlers.GetCity)
}