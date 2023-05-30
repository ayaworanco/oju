package application

import (
	"log"

	"oju/internal/parser"
	"oju/internal/tracer"
)

type Message struct {
	Type    string
	Payload interface{}
}

type Metadata struct {
	Key        string
	Host       string
	WatchQuery string
}

type Application struct {
	tree     *parser.Tree
	traces   chan []tracer.Trace
	mailbox  chan Message
	metadata Metadata
}

func Start(depth int, metadata Metadata) *Application {
	mailbox := make(chan Message)
	tree := parser.NewTree(depth)

	application := &Application{
		tree:     tree,
		mailbox:  mailbox,
		metadata: metadata,
		traces:   make(chan []tracer.Trace),
	}

	go application.run()

	return application
}

func (application *Application) SendMessage(message Message) {
	application.mailbox <- message
}

func (application *Application) run() {
	for {
		application.handle_msg(<-application.mailbox)
	}
}

func (application *Application) GetMetadata() Metadata {
	return application.metadata
}

func (application *Application) GetTraces() chan []tracer.Trace {
	return application.traces
}

func (application *Application) handle_msg(msg Message) {
	switch msg.Type {
	case "LOG":
		parser.ParseLog(application.tree, msg.Payload.(string))
	case "TRACE":
		trace, parse_error := tracer.Parse(msg.Payload.(string))
		if parse_error != nil {
			log.Println("error on building tracer", parse_error.Error())
		}
		// application.traces <- append(<-application.traces, trace)
		// to read: <-application.traces
		application.traces <- append(<-application.traces, trace)
	}
}
