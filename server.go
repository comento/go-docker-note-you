package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Define the HTTP routes
	e.File("/", "public/index.html")

	/*e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})*/

	e.Logger.Fatal(e.Start(":1323"))
}