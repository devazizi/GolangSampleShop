package router

import (
	"app/controllers"

	echo "github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", controllers.HomeAPI)
}
