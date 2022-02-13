package internal

import (
	"sync/atomic"

	"github.com/go-pkgz/lgr"
)

var status uint32

func StatusHealth(l lgr.L) {
	atomic.StoreUint32(&status, 1)
	l.Logf("[INFO] set app status health (1)")
}

func StatusNotHealth(l lgr.L) {
	atomic.StoreUint32(&status, 0)
	l.Logf("[INFO] set app status not health (0)")
}

func GetStatus() uint32 {
	return atomic.LoadUint32(&status)
}
