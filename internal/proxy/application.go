package proxy

import (
	"errors"
	"oju/internal/parser"
	"oju/internal/tracer"
)

type Metadata struct {
	Key        string
	Host       string
	WatchQuery string
}

type ApplicationMessage struct {
	Type    string
	Payload string
}

type Application struct {
	parse_tree *parser.Tree
	traces     map[string]*tracer.Trace
	metadata   Metadata
	errors     []error
}

func (application *Application) GetMetadata() Metadata {
	return application.metadata
}

func (application *Application) GetTraces() map[string]*tracer.Trace {
	return application.traces
}

func (application *Application) HandleMessage(message ApplicationMessage) {
	switch message.Type {
	case "LOG":
		parser.ParseLog(application.parse_tree, message.Payload)
	case "TRACE":
		trace, parse_trace_error := tracer.Parse(message.Payload)

		if parse_trace_error != nil {
			application.errors = append(application.errors, errors.New(parse_trace_error.Error()))
			break
		}

		application.traces[trace.GetId()] = &trace
	}
}
