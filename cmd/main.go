package main

import (
	"context"
	"github.com/bigspawn/music-news/internal"
	"github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var Lgr = lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc)

func main() {
	opt := &internal.Options{}
	p := flags.NewParser(opt, flags.Default)
	if _, err := p.Parse(); err != nil {
		Lgr.Logf("[FATAL] parse flags %v", err)
	}

	go metrics()

	Lgr.Logf("[INFO] %v", opt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := internal.NewApp(ctx, opt, Lgr)
	if err != nil {
		Lgr.Logf("[FATAL] init application: %w", err)
	}

	app.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	Lgr.Logf("[INFO] system signal %s", <-ch)

	cancel()
	app.Stop()
}

func metrics() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":9091", nil); err != nil {
		Lgr.Logf("[ERROR] metrics handler: err=%v", err)
	}
}
