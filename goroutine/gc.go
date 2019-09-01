package goroutine

import (
	"runtime"
	"runtime/debug"
	"time"
)

func GCTest() {
	test()
	runtime.GC()
	debug.FreeOSMemory()
	time.Sleep(3600 * time.Second)
}

func test() {

}