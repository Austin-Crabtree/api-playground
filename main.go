package main

import (
	"api-playground/models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
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
	e.POST("/api/upload", upload)

	// Output the current routes to a file for /api/routes to use
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Fatal(err)
	}
	err = ioutil.WriteFile("routes.json", data, 0644)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Start the server and log to stderr if a error occurs.
	// ip: 0.0.0.0 port: 1323
	e.Logger.Fatal(e.Start(":1323"))
}

// index is a simple route that sends to client 200 and Hello, World!
func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// data route creates an instance of the model Data
// fills that instance with random information and
// then sends that as json back to the client
func data(c echo.Context) error {
	clientData := models.Data{}
	clientData.FillRandom()
	return c.JSON(http.StatusOK, clientData)
}

// getOS route grabs the go runtime OS variable and
// sends it to the client
func getOS(c echo.Context) error {
	return c.String(http.StatusOK, runtime.GOOS)
}

// routes route reads the file routes.json that was
// generated on application startup and sends it to
// client. If read in the file errors out then the
// application sends 500 and an error message.
func routes(c echo.Context) error {
	routes, err := ioutil.ReadFile("routes.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading routes file: %s", err.Error()))
	}
	return c.String(http.StatusOK, string(routes))
}

func upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error getting file form value: %s", err.Error()))
	}
	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error opening file: %s", err.Error()))
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating destination file: %s", err.Error()))
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error copying file contents to destination file: %s", err.Error()))
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully</p>", file.Filename))
}
