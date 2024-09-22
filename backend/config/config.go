package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPServer HTTPServerConfig
	Database   DatabaseConfig
}

type HTTPServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewConfig() Config {
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	return Config{
		HTTPServer: HTTPServerConfig{
			Port: serverPort,
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     dbPort,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}
}
