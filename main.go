package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"golang_project/api"
	"golang_project/config"
	"golang_project/repository"
	"golang_project/server"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		fmt.Fprintf(os.Stderr, "this is the startup error: %+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, configErr := config.NewConfig()
	if configErr != nil {
		return configErr
	}

	time.Sleep(10 * time.Second)
	log.Println("sleeping for 10 seconds")
	db, err := sql.Open(cfg.GetDBDriver(), cfg.GetConnStringWithDB())
	if err != nil {
		return errors.WithStack(err)
	}

	storage, err := repository.NewStorage(db, *cfg)
	if err != nil {
		return err
	}

	//nolint:errcheck // not needed here
	defer storage.CloseDB()

	userDb := repository.NewUserStorage(db)

	userApi := api.NewUserApi(userDb)

	userHandler := server.NewUserHandler(userApi)

	srv := server.NewServer(cfg, *userHandler)

	return srv.Run()
}
