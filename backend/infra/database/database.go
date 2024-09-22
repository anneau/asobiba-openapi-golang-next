package database

import (
	"database/sql"
	"fmt"

	"github.com/anneau/asobiba-openapi-golang-next/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(config *config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	conn, err := db.DB()

	if err != nil {
		return nil, err
	}

	return conn, nil
}
