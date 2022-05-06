package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

func GetCities(c echo.Context) error {
	var result []map[string]interface{}

	url := "https://spott.p.rapidapi.com/places/autocomplete?type=CITY"

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetHeader("X-RapidAPI-Host", os.Getenv("RAPIDAPI_HOST")).
		SetHeader("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY")).
		SetResult(&result).
		Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(resp.StatusCode(), result)
}

func GetCity(c echo.Context) error {
	apiUrl := os.Getenv("WEATHERAPI_API_URL")
	apiKey := os.Getenv("WEATHERAPI_API_KEY")
	q := c.Param("q")
	url := fmt.Sprintf("%scurrent.json?key=%s&q=%s&aqi=no", apiUrl, apiKey, q)

	var result map[string]interface{}

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetResult(&result).
		Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(resp.StatusCode(), result)
}
