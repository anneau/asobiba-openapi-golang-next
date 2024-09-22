package main

import (
	"github.com/anneau/asobiba-openapi-golang-next/api/server"
	"github.com/anneau/asobiba-openapi-golang-next/config"
	"github.com/anneau/asobiba-openapi-golang-next/infra/database"
)

func main() {
	config := config.NewConfig()

	dbConn, err := database.NewConnection(&config.Database)

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	server, err := server.NewServer(&config.HTTPServer, dbConn)

	if err != nil {
		panic(err)
	}

	server.Run()
}
