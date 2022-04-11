package main

import "echo-weather/src/router"

func main() {
	// create a new echo instance
	e := router.New()
	e.Logger.Fatal(e.Start(":8000"))
}