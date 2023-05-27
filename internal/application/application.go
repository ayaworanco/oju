package application

import (
	"fmt"

	"oju/internal/parser"
)

type Message struct {
	Type    string
	Payload interface{}
}

type Application struct {
	tree        *parser.Tree
	key         string
	watch_query string
	mailbox     chan Message
}

// Starts application by running the actor model in a loop with your mailbox
func Start(tree *parser.Tree, key, watch_query string) *Application {
	mailbox := make(chan Message, 0) // maybe a []chan message?
	application := &Application{
		tree:        tree,
		key:         key,
		watch_query: watch_query,
		mailbox:     mailbox,
	}

	go application.run()

	return application
}

func (application *Application) SendMessage(msg Message) {
	application.mailbox <- msg
}

func (application *Application) run() {
	for {
		select {
		case msg := <-application.mailbox:
			application.run_msg(msg)
		}
	}
}

func (application *Application) run_msg(msg Message) {
	switch msg.Type {
	case "LOG":
		parser.ParseLog(application.tree, msg.Payload.(string))
		break
	case "WATCH":
		fmt.Println("this is a watch")
		fmt.Printf("%#v", msg.Payload)
		break
	}
}
