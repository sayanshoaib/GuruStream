package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Room(c echo.Context) error {
	return c.Render(http.StatusOK, "room", nil)
}
