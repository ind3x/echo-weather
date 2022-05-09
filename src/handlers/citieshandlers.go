package handlers

import (
	"echo-weather/src/vo"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

func GetCities(c echo.Context) error {
	var result []map[string]interface{}

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetHeader("X-RapidAPI-Host", os.Getenv("RAPIDAPI_HOST")).
		SetHeader("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY")).
		SetQueryParams(map[string]string{
			"q": c.FormValue("q"),
			"type": "CITY",
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%splaces/autocomplete", os.Getenv("RAPIDAPI_URL")))

	if err != nil {
		log.Fatalln(err)
	}

	// Map result to City
	cities := []vo.City{}
	for _, city := range result {
		cities = append(cities, vo.City{
			Name:    city["name"].(string),
			Country: city["country"].(map[string]interface{})["name"].(string),
		})
	}

	return c.JSON(resp.StatusCode(), cities)
}

func GetCity(c echo.Context) error {
	var result map[string]interface{}

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"key": os.Getenv("WEATHERAPI_KEY"),
			"q": c.Param("q"),
			"aqi": "no",
		}).
		SetResult(&result).
		Get(fmt.Sprintf("%scurrent.json", os.Getenv("WEATHERAPI_URL")))

	if err != nil {
		log.Fatalln(err)
	}

	// Map result to City
	city := vo.City{
		Name:      result["location"].(map[string]interface{})["name"].(string),
		Country:   result["location"].(map[string]interface{})["country"].(string),
		Condition: result["current"].(map[string]interface{})["condition"].(map[string]interface{})["text"].(string),
	}

	return c.JSON(resp.StatusCode(), city)
}
