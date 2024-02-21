package armazen

import (
	"fmt"
)

type Armazen struct {
	Mailbox chan Command
}

func NewArmazen() Armazen {
	armazen := Armazen{
		Mailbox: make(chan Command),
	}

	go run(armazen)
	return armazen
}

func run(armazen Armazen) {
	for {
		message := <-armazen.Mailbox
		resolve(message)
	}
}

func resolve(message Command) {
	switch message.GetType() {
	case INSERT_TRACK:
		fmt.Println("inserting a track")
	case GET_TRACK:
		fmt.Println("getting a track")
	default:
		fmt.Println("doing nothing")
	}
}
