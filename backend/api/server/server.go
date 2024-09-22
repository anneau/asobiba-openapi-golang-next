package server

import (
	"database/sql"
	"fmt"

	"github.com/anneau/asobiba-openapi-golang-next/api/router"
	"github.com/anneau/asobiba-openapi-golang-next/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type HTTPServer struct {
	engine *echo.Echo
	config *config.HTTPServerConfig
}

func NewServer(config *config.HTTPServerConfig, dbConn *sql.DB) (*HTTPServer, error) {
	engine := echo.New()

	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())

	router.Setup(*engine.Group("/api"))

	return &HTTPServer{engine: engine}, nil
}

func (s *HTTPServer) Run() {
	s.engine.Logger.Debug(s.engine.Start(fmt.Sprintf(":%d", s.config.Port)))
}
