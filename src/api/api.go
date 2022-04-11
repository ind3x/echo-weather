package api

import (
	"echo-weather/src/handlers"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)
}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}