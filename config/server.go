package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

const (
	DefaultPort = "8080"
	DefaultIp   = "0.0.0.0"
)

type serverConfig struct {
	env     string
	ip      string
	port    string
	version string
}

func newServerConfig() (*serverConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.WithStack(errors.New("config: unable to load env file"))
	}
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		return nil, errors.WithStack(errors.New("config: env was empty"))
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
	projectVersion := os.Getenv("PROJECT_VERSION")
	if env != "testing" && projectVersion == "" {
		return nil, errors.WithStack(errors.New("config: project version was empty"))
	}
	return &serverConfig{
		env:     env,
		ip:      DefaultIp,
		port:    port,
		version: projectVersion,
	}, nil
}
