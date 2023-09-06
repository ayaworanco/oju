package usecases

import (
	"fmt"
	"oju/internal/domain/entities"
	"time"
)

func new_updated_system(graph entities.Graph, mailbox chan entities.Command, resources []entities.Resource) entities.System {
	return entities.System{
		Graph:     graph,
		Resources: resources,
		Mailbox:   mailbox,
	}
}

func NewSystem(resources []entities.Resource) entities.System {
	system := entities.System{
		Graph:     new_graph(make(map[string]entities.Vertex)),
		Resources: resources,
		Mailbox:   make(chan entities.Command),
	}

	go run(system)
	return system
}

func Send(sys entities.System, message entities.Command) {
	sys.Mailbox <- message
}

func run(sys entities.System) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case message := <-sys.Mailbox:
			sys = resolve_message(message, sys)
			// save_system(sys)
		case t := <-ticker.C:
			fmt.Println("time is over: ", t)
			return
		}
	}
}

func resolve_message(message entities.Command, sys entities.System) entities.System {
	if message.GetType() == entities.INSERT_ACTION {
		return new_updated_system(update_graph(sys.Graph, message.(entities.InsertActionCommand)), sys.Mailbox, sys.Resources)
	}

	if message.GetType() == entities.GET_JOURNEY {
		message_data := message.(entities.GetJourneyCommand).Data
		message_channel := message.(entities.GetJourneyCommand).JourneyMap
		node := sys.Graph.Vertices[message_data]

		journey_map := get_journey(message_data, map[string]entities.Vertex{message_data: node}, sys.Graph.Vertices)

		message_channel <- journey_map
		return sys
	}
	return sys
}
