package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

    // respond with 200 to "/"
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<3")
	})

    // respond with 200 to "/ping"
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

    // respond with 200 to "/spock" - rewrite status code as "Live long and prosper"
	e.GET("/spock", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Live long and prosper"})
	})


	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8085"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
