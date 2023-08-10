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
	done  chan struct{}
	lgr   lgr.L
	store *Store
	ch    chan []News
}

func NewReNotifiedJob(params ReNotifiedJobParams) *ReNotifiedJob {
	return &ReNotifiedJob{
		done:  make(chan struct{}),
		lgr:   params.lgr,
		store: params.store,
		ch:    params.ch,
	}
}

func (j *ReNotifiedJob) Run(ctx context.Context) {
	j.lgr.Logf("[INFO] %s: run", j.Name())

	list, err := j.store.GetUnpublished(ctx)
	if err != nil {
		j.lgr.Logf("[ERROR] %s: failed to get unpublished posts: %v", j.Name(), err)
	}

	for _, n := range list {
		select {
		case <-ctx.Done():
			j.lgr.Logf("[INFO] %s: ctx done", j.Name())
			return
		case <-j.done:
			j.lgr.Logf("[INFO] %s: chanel done", j.Name())
			return
		default:
			j.lgr.Logf("[INFO] %s: %s", j.Name(), n.Title)
			j.ch <- []News{n}
		}
	}

	j.lgr.Logf("[INFO] %s: finish, published %d posts", j.Name(), len(list))
}

func (j *ReNotifiedJob) Stop() {
	close(j.done)
}

func (j *ReNotifiedJob) Name() string {
	return "re-notified-job"
}
