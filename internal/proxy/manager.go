package proxy

import (
	"oju/internal/config"
	"oju/internal/parser"
	"oju/internal/tracer"
)

const ACTION_GET_TRACES = "GET_TRACES"

type Manager struct {
	Applications []*Application
	Mailbox      chan Message
}

type Message struct {
	Destination string
	Payload     ApplicationMessage
}

func NewManager(allowed_applications []config.Application) *Manager {
	manager := &Manager{
		Applications: make([]*Application, 0),
		Mailbox:      make(chan Message),
	}

	for _, allowed := range allowed_applications {
		manager.Applications = append(manager.Applications, get_app(allowed))
	}

	return manager
}

func (manager *Manager) Redirect(destination string, payload ApplicationMessage) {
	message := Message{
		Destination: destination,
		Payload:     payload,
	}

	for _, app := range manager.Applications {
		metadata := app.GetMetadata()
		if metadata.Host == message.Destination || metadata.Key == message.Destination {
			app.HandleMessage(message.Payload)
		}
	}
}

func (proxy *Manager) GetAppTraces(destination string) map[string]*tracer.Trace {
	for _, app := range proxy.Applications {
		metadata := app.GetMetadata()
		if metadata.Host == destination || metadata.Key == destination {
			return app.traces
		}
	}
	return nil
}

func get_app(config_app config.Application) *Application {
	return &Application{
		parse_tree: parser.NewTree(10),
		traces:     make(map[string]*tracer.Trace),
		metadata: Metadata{
			Key:  config_app.AppKey,
			Host: config_app.Host,
		},
	}
}
