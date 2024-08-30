package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5/middleware"

	"api/infrastructure/config"
	"api/internal/core"
	transportHttp "api/internal/interfaces/web/router"
	"api/internal/repo/sqllite"
	"api/pkg/servers/static"
)

type App struct {
	httpServer *http.Server

	recordCore *core.Core
}

func New(cnf *config.Config, logger *zap.Logger) *App {
	db, err := sqllite.New(cnf.Db.ConnStr)
	if err != nil {
		panic(err.Error())
	}

	recordCore, err := core.New(logger, cnf, db)
	if err != nil {
		panic(err.Error()) // TODO !!!
	}

	return &App{
		recordCore: recordCore,
	}
}

func (a *App) Run(apiConfig *config.Api) error {
	router := chi.NewRouter()

	staticServer := static.New(apiConfig.Static.FilesPath)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(staticServer.Handler)

	r := transportHttp.RegisterHTTPEndpoints(router, *a.recordCore, apiConfig)

	a.httpServer = &http.Server{
		Addr:           apiConfig.Host + ":" + apiConfig.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if apiConfig.Tls.Enable {
		a.startTls(apiConfig)
	} else {
		a.startWithoutTls()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func (a *App) startTls(cfg *config.Api) {
	go func() {
		if err := a.httpServer.ListenAndServeTLS(
			cfg.Tls.CertFilePath,
			cfg.Tls.KeyFilePath,
		); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
}

func (a *App) startWithoutTls() {
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
}
