package internal

import (
	"net/http"

	"github.com/go-pkgz/lgr"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func health(l lgr.L) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if getStatus(l) == 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

func StartHandlers(l lgr.L) {
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/healthz", health(l))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		l.Logf("[ERROR] metrics handler: err=%v", err)
	}
}
