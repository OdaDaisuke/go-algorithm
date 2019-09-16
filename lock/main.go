package lock

import (
	"sync"
	"fmt"
)

func Start() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	var counter uint16
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func withAtomis() {
	//atomic.AddInt32()
}