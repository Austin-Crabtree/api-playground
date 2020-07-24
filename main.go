package main

import (
	"api-playground/models"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main()  {
	// Create an Echo instance
	e := echo.New()

	// Setup middleware to use
	// Logger will log http requests to stdout. Here it is configured
	// to output only a subset of information
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	// Recover middleware recovers from panics anywhere in the chain,
	// prints stack trace and handles the control to the centralized HTTPErrorHandler.
	e.Use(middleware.Recover())

	// Setup routes for the application
	e.GET("/", index)
	e.GET("/api/data", data)
	e.GET("/api/os", getOS)
	e.GET("/api/routes", routes)

	// Output the current routes to a file for /api/routes to use

	// Start the server and log to stderr if a error occurs.
	// ip: 0.0.0.0 port: 1323
	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func data(c echo.Context) error {
	clientData := models.Data{}
	clientData.FillRandom()
	return c.JSON(http.StatusOK, clientData)
}

func getOS(c echo.Context) error {
	return c.String(http.StatusOK, runtime.GOOS)
}

func routes(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, )
}
