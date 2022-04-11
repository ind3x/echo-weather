package handlers

import (
	"echo-weather/src/vo"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	//requestID := uuid.NewV4()
	//c.Logger().Infof("RequestID: %s", requestID)

	resp := vo.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}