package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/automaxprocs/maxprocs"

	"github.com/bigspawn/music-news/internal"
)

func main() {
	logger := lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc)

	_, err := maxprocs.Set(maxprocs.Logger(logger.Logf))
	if err != nil {
		logger.Logf("[FATAL] cant sent max proc err=%v", err)
	}

	opt := &internal.Options{}
	p := flags.NewParser(opt, flags.Default)
	if _, err := p.Parse(); err != nil {
		logger.Logf("[FATAL] parse flags err=%v", err)
	}

	go internal.StartHandlers(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, err := internal.NewApp(ctx, opt, logger)
	if err != nil {
		logger.Logf("[FATAL] init application: err=%v", err)
	}

	err = app.Start(ctx)
	if err != nil {
		logger.Logf("[FATAL] start application: err=%v", err)
	}

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	logger.Logf("[INFO] system signal %s", <-ch)

	cancel()
	app.Stop()
}
