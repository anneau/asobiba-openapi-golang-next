package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

type HealthDTO struct {
	Status string `json:"status"`
}

func (c *HealthController) Check(ctx echo.Context) error {
	res := HealthDTO{
		Status: "ok",
	}
	return ctx.JSON(http.StatusOK, res)
}
