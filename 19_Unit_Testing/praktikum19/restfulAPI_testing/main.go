package main

import (
	"testing/restfulAPI_testing/config"
	"testing/restfulAPI_testing/route"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	config.InitDB()
	e := echo.New()
	route.Route(e)
	// Route / to handler function

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
