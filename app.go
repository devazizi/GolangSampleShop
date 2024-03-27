package main

import (
	"app/router"

	// "reflect"
	// "fmt"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.Routes(e)
	// fmt.Println("hello world this is a simple golang application")
	e.Logger.Fatal(e.Start("0.0.0.0:3400"))
}
