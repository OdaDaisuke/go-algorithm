package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func Serial(num []int) {
	for _, n := range num {
		fmt.Printf("serial: %d\n", n)
	}
}

func Parallel(num []int) {
	fmt.Println()
	var wp sync.WaitGroup

	for _, n := range num {
		wp.Add(1)

		go func(n int) {
			fmt.Printf("parallel: %d\n", n)
			wp.Done()
		}(n)

	}

	wp.Wait()
}

func ParallelChannel(num []int) {
	res := make(chan int, len(num))
	go func() {
		defer close(res)
		for _, n := range num {
			fmt.Printf("parallel channel push: %d\n", n)
			res <- n
			time.Sleep(time.Millisecond)
		}
	}()

	for r := range res {
		fmt.Printf("parallel channel pull: %d\n", r)
	}
}