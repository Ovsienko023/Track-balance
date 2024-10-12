package main

import (
	"api/gen/api/apiconnect"
	"api/internal/interfaces"
	"embed"
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"

	"api/infrastructure/config"
	"api/infrastructure/logger"
)

// TODO list:
//
// [] infrastructure
//     [] config
//     [] logger
//     [] linter
//     [] make
//     [] Исследовать библиотеки для валидации
//     [x] Исследовать библиотеки для генерации документации
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
	lgr.Info(fmt.Sprintf("View API documentation: http://%s:%s/web/doc \n", cfg.Api.Host, cfg.Api.Port))

	server, err := interfaces.New(lgr, cfg)
	if err != nil {
		log.Fatalf("Could not start server with error: %+v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/web/", http.StripPrefix("/", http.FileServer(http.FS(fs))))

	path, handler := apiconnect.NewServiceHandler(server)

	mux.Handle(path, handler)
	http.ListenAndServe(
		fmt.Sprintf("%s:%s", cfg.Api.Host, cfg.Api.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

	//app := server.New(cfg, lgr, &fs)
	//if err := app.Run(&cfg.Api); err != nil {
	//	log.Fatalf("%s", err.Error())
	//}
}
