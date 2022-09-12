package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MentorDashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "mentorDashboard", nil)
}
