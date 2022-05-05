package main

import (
	"echo-weather/src/router"

	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	godotenv.Load(".env")

	// create a new echo instance
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}