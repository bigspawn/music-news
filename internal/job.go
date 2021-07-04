package internal

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/go-pkgz/lgr"
)

func NewJob(s MusicScraper, sch *gocron.Scheduler, name string, lgr lgr.L) *Job {
	return &Job{
		s:    s,
		sch:  sch,
		name: name,
		lgr:  lgr,
	}
}

type Job struct {
	s    MusicScraper
	sch  *gocron.Scheduler
	name string
	lgr  lgr.L
}

func (j Job) Do(ctx context.Context) {
	if err := j.s.Scrape(ctx); err != nil {
		j.lgr.Logf("[ERROR] %s scraper %v", j.name, err)
	}

	_, next := j.sch.NextRun()
	j.lgr.Logf("[INFO] %s job next start %s", j.name, next)
}
