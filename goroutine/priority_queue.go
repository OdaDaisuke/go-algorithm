package goroutine

import (
	"fmt"
	"time"
)

type PriorityChannel struct {
	High chan interface{}
	Normal chan interface{}
	Low chan interface{}

	Out chan interface{}
	stopCh chan struct{}
}

func NewPriorityChannel() *PriorityChannel {
	pc := PriorityChannel{}
	pc.Out = make(chan interface{})
	pc.High = make(chan interface{})
	pc.Normal = make(chan interface{})
	pc.Low = make(chan interface{})
	pc.stopCh = make(chan struct{})

	pc.start()
	return &pc
}

func (pc *PriorityChannel) start() {
	go func() {
		for {
			// High
			select {
			case s := <- pc.High:
				pc.Out <- s
				continue
			case <- pc.stopCh:
				return
			default:
			}

			// Normal
			select {
			case s := <- pc.High:
				pc.Out <- s
				continue
			case s := <- pc.Normal:
				pc.Out <- s
				continue
			case <- pc.stopCh:
				return
			default:
			}

			// Low
			select {
			case s := <- pc.High:
				pc.Out <- s
			case s := <- pc.Normal:
				pc.Out <- s
			case s := <- pc.Low:
				pc.Out <- s
			case <- pc.stopCh:
				return
			}
		}
	}()
}

func startConsumer(queue *PriorityChannel) {
	for {
		select {
		case t := <- queue.Out:
			fmt.Printf("consumed %s\n", t)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func StartPriorityQueue() {
	var queue = NewPriorityChannel()
	for i := 0; i < 4; i++ {
		go func(i int) {
			queue.Normal <- fmt.Sprintf("normal: %d", i)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 2; i++ {
		go func(i int) {
			queue.High <- fmt.Sprintf("high: %d", i)
		}(i)
	}

	go startConsumer(queue)

	// 全goroutineが待機状態になったら落ちる
	select {}
}
