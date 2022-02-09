package internal

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/go-pkgz/lgr"
)

type Job struct {
	s    MusicScraper
	sch  *gocron.Scheduler
	name string
	lgr  lgr.L
}

func (j Job) Do(ctx context.Context) {
	if err := j.s.Scrape(ctx); err != nil {
		StatusNotHealth()

		j.lgr.Logf("[ERROR] %s scraper %v", j.name, err)
	}

	_, next := j.sch.NextRun()
	j.lgr.Logf("[INFO] %s job next start %s", j.name, next)
}
