package proxy

import (
	"errors"
	"oju/internal/tracer"
)

type StackTrace struct {
	traces []*tracer.Trace
	errors []error
}

func NewStackTrace() *StackTrace {
	return &StackTrace{
		traces: make([]*tracer.Trace, 0),
	}
}

func (stack *StackTrace) RunStack(trace *tracer.Trace, metadatas []Metadata) {
	if len(stack.traces) == 0 {
		stack.traces = append(stack.traces, trace)
	} else {
		stack.traces = append(stack.traces, trace)
		previous_index := len(stack.traces) - 2
		previous_trace := stack.traces[previous_index]

		if previous_trace.Service == trace.AppKey {
			previous_trace.AddChild(trace)

			stack.traces = remove_trace(stack.traces, previous_index)
			return
		}

		host, host_error := get_host_by_app_key(trace.AppKey, metadatas)
		if host_error != nil {
			stack.errors = append(stack.errors, host_error)
			return
		}

		if previous_trace.Service == host {
			previous_trace.AddChild(trace)

			stack.traces = remove_trace(stack.traces, previous_index)
			return
		}

		if previous_trace.Service == "" {
			for _, value := range previous_trace.Attributes {
				if value == trace.AppKey || value == host {
					previous_trace.AddChild(trace)
				}
			}
		}

		stack.traces = remove_trace(stack.traces, previous_index)
	}
}

func get_host_by_app_key(app_key string, metadatas []Metadata) (string, error) {
	for _, metadata := range metadatas {
		if metadata.Key == app_key {
			return metadata.Host, nil
		}
	}

	return "", errors.New("host not found")
}

func remove_trace(traces []*tracer.Trace, index int) []*tracer.Trace {
	return append(traces[:index], traces[index:]...)
}
