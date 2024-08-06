package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type databaseConfig struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbName     string
	dbDriver   string
	dbPort     string
}

func newDatabaseConfig(env string) (*databaseConfig, error) {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return nil, errors.WithStack(errors.New("storage: dbUser env variable"))
	}
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		return nil, errors.WithStack(errors.New("storage: dbPass env variable"))
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return nil, errors.WithStack(errors.New("storage: dbHost env variable"))
	}
	dbName := os.Getenv("DB_NAME")
	if env != "testing" && dbName == "" {
		return nil, errors.WithStack(errors.New("storage: dbName env variable"))
	}
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		return nil, errors.WithStack(errors.New("storage: dbDriver env variable"))
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return nil, errors.WithStack(errors.New("storage: dbPort env variable"))
	}
	return &databaseConfig{
		dbUser:     dbUser,
		dbPassword: dbPass,
		dbHost:     dbHost,
		dbName:     dbName,
		dbDriver:   dbDriver,
		dbPort:     dbPort,
	}, nil
}

func (d *databaseConfig) setupConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true", d.dbUser, d.dbPassword, d.dbHost, d.dbName)
}

// func (d *databaseConfig) setupConnectionString() string {
// 	connectionStr := fmt.Sprintf(
// 		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
// 		d.dbUser, d.dbPassword, d.dbHost, d.dbPort, d.dbName,
// 	)
// 	return connectionStr
// }
