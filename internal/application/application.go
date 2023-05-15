package application

import (
	"fmt"
	"oluwoye/internal/parser"
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

/*
This represents an application with your own parse tree
app key and watch query (that I will try to update to more than one queries to be watching)
*/
func StartApplication(tree *parser.Tree, key, watch_query string) *Application {
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
			run_msg(application, msg)
		}
	}
}

func run_msg(application *Application, msg Message) {
	switch msg.Type {
	case "LOG":
		// TODO: need to do the parse correctly
		// this Id need to be always new like an uuid
		parser.ParseLog(application.tree, msg.Payload.(string), 1) // FIXME: add an uuid here as a unique
		break
	case "WATCH":
		fmt.Println("this is a watch")
		fmt.Printf("%#v", msg.Payload)
		break
	}
}
