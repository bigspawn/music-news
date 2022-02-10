package main

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/bigspawn/music-news/internal"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	logger := lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc)

	opt := &internal.Options{}
	p := flags.NewParser(opt, flags.Default)
	if _, err := p.Parse(); err != nil {
		logger.Logf("[FATAL] parse flags err=%v", err)
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.Handle("/health", health())

		if err := http.ListenAndServe(":8080", nil); err != nil {
			logger.Logf("[ERROR] metrics handler: err=%v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := internal.NewApp(ctx, opt, logger)
	if err != nil {
		logger.Logf("[FATAL] init application: err=%v", err)
	}

	go app.Start()

	internal.StatusHealth()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	logger.Logf("[INFO] system signal %s", <-ch)

	cancel()
	app.Stop()
}

func health() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		v := internal.GetStatus()
		if v == 0 {
			writer.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		writer.WriteHeader(http.StatusOK)
	})
}
