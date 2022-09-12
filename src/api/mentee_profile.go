package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MenteeProfile(c echo.Context) error {
	return c.Render(http.StatusOK, "menteeProfile", nil)
}
