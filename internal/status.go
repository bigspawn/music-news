package internal

import (
	"sync/atomic"

	"github.com/go-pkgz/lgr"
)

var status uint32

func getStatus(l lgr.L) uint32 {
	val := atomic.LoadUint32(&status)
	return val
}

func StatusHealth(l lgr.L) {
	atomic.StoreUint32(&status, 1)
	l.Logf("[INFO] set app status health (1)")
}
