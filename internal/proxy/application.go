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

func (application *Application) HandleMessage(message ApplicationMessage, stack_trace *StackTrace, applications_metadata []Metadata) {
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
		service, service_error := get_service(trace, applications_metadata)

		if service_error != nil {
			application.errors = append(application.errors, service_error)
			break
		}

		stack_trace.RunStack(&trace, service)
	}
}

func get_service(trace tracer.Trace, metadatas []Metadata) (string, error) {
	service := trace.Service
	attributes := trace.Attributes
	for _, metadata := range metadatas {
		if service != "" {
			if metadata.Key == service {
				return metadata.Key, nil
			}
			if metadata.Host == service {
				return metadata.Host, nil
			}
		} else {
			for _, value := range attributes {
				if value == metadata.Key || value == metadata.Host  {
					return value, nil
				}
			}
		}
	}

	return "", errors.New("service not found")
}
