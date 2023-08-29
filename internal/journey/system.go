package journey

import (
	"oju/internal/config"
	//"oju/internal/tracer"
	//"fmt"
)

const (
	INSERT_ACTION  = "insert_action"
	GET_STRUCTURE = "get_structure"
)

type system struct {
	graph   graph
	applications []config.Application
	mailbox chan interface{} 
}


func NewSystem(allowed_applications []config.Application) system {
	system := system{
		graph:   new_graph(make(map[string]vertex)),
		applications: allowed_applications,
		mailbox: make(chan interface{}),
	}

	go run(system)
	return system
}

func Send(sys system, message interface{}) {
	sys.mailbox <- message
}

func run(sys system) {
	for message := range sys.mailbox {
		println(message)
		/*
		switch message.Type {
		case INSERT_ACTION:
			//system.graph = update_graph[T](system.graph, message)
			fmt.Println("inserting action")
			break
		case GET_STRUCTURE:
			fmt.Println("getting structure")
			break
		}
		*/
	}
}
