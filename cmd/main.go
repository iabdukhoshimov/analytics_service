package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"gitlab.com/greatsoft/xif-backend/internal/config"
	"gitlab.com/greatsoft/xif-backend/internal/pkg/logger"
	"gitlab.com/greatsoft/xif-backend/internal/transport/handlers"
	"gitlab.com/greatsoft/xif-backend/pkg/logger/factory"

	autometrics "github.com/autometrics-dev/autometrics-go/pkg/autometrics/otel"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cfg *config.Config
)

func init() {
	log.Println("Initializing config...")
	cfg = config.Load()

	customLog, err := factory.Build(&cfg.Logging)
	if err != nil {
		log.Fatal(err)
	}

	logger.SetLogger(customLog)
}

func main() {
	server := handlers.NewServer(cfg)

	go func() {
		server.Run()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	gracefulShutdown(server, ctx, cancel)

	autometrics.Init(
		"myApp/v2/prod",
		autometrics.DefBuckets,
		autometrics.BuildInfo{
			Version: "0.4.0",
			Commit:  "anySHA",
			Branch:  "",
		},
	)
	http.Handle("/metrics", promhttp.Handler())
}

func gracefulShutdown(server handlers.Server, ctx context.Context, cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		logger.Log.Info("shutting down")
		server.Stop()
		logger.Log.Info("shutdown successfully called")
		wg.Done()
	}(&wg)

	go func() {
		wg.Wait()
		cancel()
	}()

	<-ctx.Done()
	os.Exit(0)
}
