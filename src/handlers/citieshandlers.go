package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func GetCities(c echo.Context) error {
	return c.String(http.StatusOK, " You are on the Admin Page !!!")
}

func GetCity(c echo.Context) error {
	q := c.Param("q")
	apiUrl := os.Getenv("WEATHERAPI_API_URL")
	apiKey := os.Getenv("WEATHERAPI_API_KEY")

	resp, err := http.Get(apiUrl + "current.json?key=" + apiKey + "&q=" + q + "&aqi=no")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	errResult := json.NewDecoder(resp.Body).Decode(&result)
	if errResult != nil {
		log.Fatalln(errResult)
	}

	return c.JSON(http.StatusOK, result)
}
