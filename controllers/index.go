package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeAPI(c echo.Context) error {
	return c.String(http.StatusBadRequest, "Request invalid")
}
