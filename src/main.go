package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"softwareLab/src/api"
	"softwareLab/src/model"
	"softwareLab/src/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	model.DbConnection()
	api.TempRender(e)
	routes.SetupRoutes(e)

	if err := e.Start(":8082"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
