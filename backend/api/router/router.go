package router

import (
	"github.com/anneau/asobiba-openapi-golang-next/api/controller"
	"github.com/labstack/echo"
)

type Router struct {
	HealthController *controller.HealthController
}

func Setup(router echo.Group) {
	health := controller.NewHealthController()
	router.GET("/health", health.Check)
}
