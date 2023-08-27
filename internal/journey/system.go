package journey

import (
	"oju/internal/config"
	"oju/internal/tracer"
	"reflect"
	"fmt"
)

const (
	INSERT_ACTION  = "insert_action"
	GET_STRUCTURE = "get_structure"
)

type system struct {
	graph   graph
	mailbox chan interface{} // FIXME: change to a better struct
}

type SystemMessage struct {
	Trace       tracer.Trace
}

func NewSystem(allowed_applications []config.Application) system {
	system := system{
		graph:   new_graph(allowed_applications),
		mailbox: make(chan interface{}),
	}

	go system.run()
	return system
}

func (system system) InsertAction(data tracer.Trace) {
	system.mailbox <- data
}

func (system system) run() {
	for message := range system.mailbox {
		message_type := reflect.TypeOf(message)
		switch message_type.Name() {
		case "Trace":
			system.graph = update_graph(system.graph, message.(tracer.Trace))
			break
		}
	}
}
