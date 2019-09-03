package goroutine

import (
	"fmt"
	"time"
)

type Origin struct {
	Receiver chan Message
}

func newOrigin() *Origin {
	return &Origin{
		Receiver: make(chan Message),
	}
}

type Message struct {
	Body   string
	RoomID int
	UserID int
}

func StartChatRoom() {
	origin := newOrigin()

	go func() {
		time.Sleep(1 * time.Second)
		origin.Receiver <- Message{
			Body:   "hello from 1 user",
			RoomID: 2,
			UserID: 1,
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		origin.Receiver <- Message{
			Body:   "hello from 2 user",
			RoomID: 2,
			UserID: 2,
		}
	}()

	go func() {
		time.Sleep(1250 * time.Millisecond)
		origin.Receiver <- Message{
			Body:   "hello from 3 user",
			RoomID: 2,
			UserID: 3,
		}
	}()

	for i := 0; i < 3; i++ {
		m := <-origin.Receiver
		fmt.Printf("msg received %v\n", m)
	}
}
