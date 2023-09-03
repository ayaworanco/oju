package system

import (
	"fmt"
	"oju/internal/config"
	"oju/internal/journey"
	"time"
)

type System struct {
	Graph        journey.Graph
	Applications []config.Application
	Mailbox      chan journey.Command
}

func new_updated_system(graph journey.Graph, mailbox chan journey.Command, applications []config.Application) System {
	return System{
		Graph:        graph,
		Applications: applications,
		Mailbox:      mailbox,
	}
}

func NewSystem(allowed_applications []config.Application) System {
	system := System{
		Graph:        journey.NewGraph(make(map[string]journey.Vertex)),
		Applications: allowed_applications,
		Mailbox:      make(chan journey.Command),
	}

	go run(system)
	return system
}

func Send(sys System, message journey.Command) {
	sys.Mailbox <- message
}

func run(sys System) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case message := <-sys.Mailbox:
			sys = resolve_message(message, sys)
		case t := <-ticker.C:
			fmt.Println("time is over: ", t)
			return
		}
	}
}

func resolve_message(message journey.Command, sys System) System {
	if message.GetType() == journey.INSERT_ACTION {
		return new_updated_system(journey.UpdateGraph(sys.Graph, message.(journey.InsertActionCommand)), sys.Mailbox, sys.Applications)
	}

	if message.GetType() == journey.GET_JOURNEY {
		message_data := message.(journey.GetJourneyCommand).Data
		message_channel := message.(journey.GetJourneyCommand).JourneyMap
		node := sys.Graph.Vertices[message_data]

		journey_map := journey.GetJourney(message_data, map[string]journey.Vertex{message_data: node}, sys.Graph.Vertices)

		message_channel <- journey_map
		return sys
	}
	return sys
}
