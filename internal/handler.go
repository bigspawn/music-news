package internal

import (
	"net/http"

	"github.com/go-pkgz/lgr"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartHandlers(l lgr.L) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		l.Logf("[ERROR] metrics handler: err=%v", err)
	}
}
