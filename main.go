package main

import (
	"embed"
	"flag"
	"fmt"
	"log"

	"api/infrastructure/config"
	"api/infrastructure/logger"
	server "api/pkg/servers"
)

// TODO list:
//
// [] infrastructure
//     [] config
//     [] logger
//     [] linter
//     [] make
//     [] Исследовать библиотеки для валидации
//     [] Исследовать библиотеки для генерации документации
// [] DB service
// [] Core
//     [] Auth

// @title Track Balance API
// @version 0.0.1
// @description This is a balance tracker server.

var (
	//go:embed web
	fs embed.FS
)

// @BasePath /api/v1
func main() {
	flag.Parse()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Could not read configuration file with error: %+v", err)
	}

	lgr, err := logger.New()
	if err != nil {
		log.Fatalf("Could not read configuration file with error: %+v", err)
	}

	lgr.Info(fmt.Sprintf("Running on: %s:%s", cfg.Api.Host, cfg.Api.Port))
	lgr.Info(fmt.Sprintf("View API documentation: http://%s:%s/api/v1/docs \n", cfg.Api.Host, cfg.Api.Port))

	app := server.New(cfg, lgr, &fs)
	if err := app.Run(&cfg.Api); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
