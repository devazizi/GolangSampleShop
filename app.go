package main

import (
	// "fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a simple Golang application")
	})

	// fmt.Println("hello world this is a simple golang application")
	e.Logger.Fatal(e.Start("0.0.0.0:3400"))
}
