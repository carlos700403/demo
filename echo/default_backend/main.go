package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func hello(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotFound, "Not Found")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/*", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
