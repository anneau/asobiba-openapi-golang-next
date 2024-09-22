package main

import (
	"net/http"

	"github.com/anneau/asobiba-openapi-golang-next/config"
	"github.com/anneau/asobiba-openapi-golang-next/infra/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config := config.NewConfig()

	dbConn, err := database.NewConnection(&config.Database)

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", health)

	e.Logger.Fatal(e.Start(":1323"))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
