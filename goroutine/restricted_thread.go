package goroutine

import (
	"fmt"
	"time"
)

func StartRestrictionThread() {
	const (
		total = 6
		maxConcurrency = 3
	)
	parentSigQueue := make(chan string, maxConcurrency)
	queueReceiver := make(chan string, total)
	for i := 0; i < total; i++ {
		go func() {
			parentSigQueue <- fmt.Sprintf("parent sig %d", i)
			time.Sleep(1 * time.Second)
			fmt.Printf("%d:end wait 1sec\n", i)
			queueReceiver <- fmt.Sprintf("child sig %d", i)
			<- parentSigQueue
		}()
	}

	for {
		if len(queueReceiver) >= total {
			break
		}
	}

	fmt.Printf("End goroutine %s \n", time.Now())
}
