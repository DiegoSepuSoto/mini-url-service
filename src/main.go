package main

import (
	"context"
	"github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/http/handlers/metrics"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/DiegoSepuSoto/mini-url-service/src/application/usecase/miniurl"
	"github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/database/repositories/mongodb/miniurls"
	redisMiniUrls "github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/database/repositories/redis/miniurls"
	"github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/http/handlers/health"
	miniurlHandler "github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/http/handlers/miniurl"
	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
)

const closeAppTimeout = time.Second * 10

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	metrics.NewMetricsHandler(e)

	health.NewHealthHandler(e)

	miniURLHandler(e)

	quit := make(chan os.Signal, 1)
	go startServer(e, quit)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	gracefulShutdown(e)
}

func miniURLHandler(e *echo.Echo) {
	mongoDBCollection := shared.CreateMongoDBCollection()
	redisClient := shared.CreateRedisClient()

	mongoDBMiniURLRepo := miniurls.NewMongoDBMiniURLsRepository(mongoDBCollection)
	redisMiniURLRepo := redisMiniUrls.NewRedisMiniURLsRepository(mongoDBMiniURLRepo, redisClient)

	miniURLUseCase := miniurl.NewMiniURLUseCase(redisMiniURLRepo)

	_ = miniurlHandler.NewMiniURLHandler(e, miniURLUseCase)
}

func startServer(e *echo.Echo, quit chan os.Signal) {
	log.Print("starting server")

	if err := e.Start(":" + os.Getenv("APP_PORT")); err != nil {
		log.Error(err.Error())
		close(quit)
	}
}

func gracefulShutdown(e *echo.Echo) {
	log.Print("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), closeAppTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
