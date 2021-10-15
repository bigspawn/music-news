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

	go metrics(logger)

	logger.Logf("[INFO] %v", opt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := internal.NewApp(ctx, opt, logger)
	if err != nil {
		logger.Logf("[FATAL] init application: err=%v", err)
	}

	app.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	logger.Logf("[INFO] system signal %s", <-ch)

	cancel()
	app.Stop()
}

func metrics(logger *lgr.Logger) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":9091", nil); err != nil {
		logger.Logf("[ERROR] metrics handler: err=%v", err)
	}
}
