package server

import (
	"context"
	"embed"
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
	"api/internal/repov2"
	"api/pkg/servers/static"
)

type App struct {
	httpServer *http.Server
	fs         *embed.FS

	recordCore *core.Core
}

func New(cnf *config.Config, logger *zap.Logger, fs *embed.FS) *App {
	//db, err := sqllite.New(cnf.Db.ConnStr)
	//if err != nil {
	//	panic(err.Error())
	//}

	driver, err := repov2.Conn(cnf.Db.ConnStr)
	if err != nil {
		panic(err.Error())
	}

	if err = repov2.InitDb(driver); err != nil {
		panic(err.Error())
	}

	recordCore, err := core.New(logger, cnf, core.Repositories{
		Users:   repov2.NewUsers(driver),
		Circles: repov2.NewCircles(driver),
		Areas:   repov2.NewAreas(driver),
	})

	if err != nil {
		panic(err.Error()) // TODO !!!
	}

	return &App{
		recordCore: recordCore,
		fs:         fs,
	}
}

func (a *App) Run(apiConfig *config.Api) error {
	router := chi.NewRouter()

	staticServer := static.New(apiConfig.Static.FilesPath)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(staticServer.Handler)
	router.Use(transportHttp.EnableCors)

	r := transportHttp.RegisterHTTPEndpoints(router, *a.recordCore, a.fs)

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
