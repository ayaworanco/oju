package system

import (
	"fmt"
	"oju/internal/config"
	"oju/internal/journey"
	"time"
)

type System struct {
	graph        journey.Graph
	applications []config.Application
	mailbox      chan journey.Command
}

func new_updated_system(graph journey.Graph, mailbox chan journey.Command, applications []config.Application) System {
	return System{
		graph:        graph,
		applications: applications,
		mailbox:      mailbox,
	}
}

func NewSystem(allowed_applications []config.Application) System {
	system := System{
		graph:        journey.NewGraph(make(map[string]journey.Vertex)),
		applications: allowed_applications,
		mailbox:      make(chan journey.Command),
	}

	// TODO: allow load by armazen the system with graph
	go run(system)
	return system
}

func Send(sys System, message journey.Command) {
	sys.mailbox <- message
}

func run(sys System) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case message := <-sys.mailbox:
			// TODO: use armazen to save when system resolve a message (check later the saving strategy)
			sys = resolve_message(message, sys)
		case t := <-ticker.C:
			fmt.Println("time is over: ", t)
			return
		}
	}
}

func resolve_message(message journey.Command, sys System) System {
	if message.GetType() == journey.INSERT_ACTION {
		return new_updated_system(journey.UpdateGraph(sys.graph, message.(journey.InsertActionCommand)), sys.mailbox, sys.applications)
	}

	if message.GetType() == journey.GET_JOURNEY {
		message_data := message.(journey.GetJourneyCommand).Data
		message_channel := message.(journey.GetJourneyCommand).JourneyMap
		node := sys.graph.Vertices[message_data]

		journey_map := journey.GetJourney(message_data, map[string]journey.Vertex{message_data: node}, sys.graph.Vertices)

		message_channel <- journey_map
		return sys
	}
	return sys
}

func AddError(system System, some_error error) {
	// TODO: use armazen to log this error
}
