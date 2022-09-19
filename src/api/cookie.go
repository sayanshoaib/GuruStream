package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ReadAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all the cookies")
}
