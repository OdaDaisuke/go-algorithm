package sync

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
	"time"
)

func Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	withNoAtomic()
	withAtomic()
	syncOnce()
	syncCond()
}

func withNoAtomic() {
	var v int32
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			v++
			v++
			v++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}

func withAtomic() {
	var v int32
	wg := new(sync.WaitGroup)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			atomic.AddInt32(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}

func syncOnce() {
	var once = new(sync.Once)

	greeting := func(wg *sync.WaitGroup) {
		once.Do(func() {
			fmt.Println("hello from once")
		})
		fmt.Println("hello!")
		wg.Done()
	}

	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go greeting(wg)
	}

	wg.Wait()
}

func syncCond() {
	// goroutineをあらかじめためておいて、後からcondで実行
	m := new(sync.Mutex)
	c := sync.NewCond(m)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("waiting %d\n", i)
			m.Lock()
			defer m.Unlock()
			c.Wait()
			fmt.Printf("num %d\n", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		c.Signal()
	}

	// ためたジョブを一気に実行
	//c.Broadcast()
}
