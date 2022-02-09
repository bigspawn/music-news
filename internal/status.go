package internal

import "sync/atomic"

var status uint32

func StatusHealth() {
	atomic.StoreUint32(&status, 1)
}

func StatusNotHealth() {
	atomic.StoreUint32(&status, 0)
}

func GetStatus() uint32 {
	return atomic.LoadUint32(&status)
}
