package internal

import (
	"context"

	"github.com/go-pkgz/lgr"
)

type ReNotifiedJobParams struct {
	lgr   lgr.L
	store *Store
	ch    chan []News
}

type ReNotifiedJob struct {
	ReNotifiedJobParams

	done chan struct{}
}

func NewReNotifiedJob(params ReNotifiedJobParams) *ReNotifiedJob {
	return &ReNotifiedJob{
		ReNotifiedJobParams: params,
		done:                make(chan struct{}),
	}
}

func (j *ReNotifiedJob) Run(ctx context.Context) {
	j.lgr.Logf("[INFO] %s: run", j.Name())

	list, err := j.store.GetUnpublished(ctx)
	if err != nil {
		j.lgr.Logf("[ERROR] %s: failed to get unpublished posts: %v", j.Name(), err)
	}

	select {
	case <-ctx.Done():
		j.lgr.Logf("[INFO] %s: ctx done", j.Name())
		return
	case <-j.done:
		j.lgr.Logf("[INFO] %s: chanel done", j.Name())
		return
	case j.ch <- list:
		j.lgr.Logf("[INFO] %s: published %d posts", j.Name(), len(list))
		return
	}
}

func (j *ReNotifiedJob) Stop() {
	close(j.done)
}

func (j *ReNotifiedJob) Name() string {
	return "re-notified-job"
}
