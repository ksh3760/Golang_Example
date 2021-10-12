package main

import (
	"github.com/labstack/echo/v4"
)

var (
	gPort string = "8080"
)

func main() {
	sE := echo.New()
	sE.File("/", "static/main.html")
	sE.Static("/", "static")
	sE.Logger.Fatal(sE.Start(":" + gPort))

	// e.GET("/", func(c echo.Context) error {
    //     return c.String(http.StatusOK, "Hello World!")
    // })
}
