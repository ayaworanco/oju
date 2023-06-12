package proxy

import (
	"fmt"
	"oju/internal/tracer"
)

type StackTrace struct {
	traces []*tracer.Trace
}

func NewStackTrace() *StackTrace {
	return &StackTrace{
		traces: make([]*tracer.Trace, 0),
	}
}

func (stack *StackTrace) RunStack(trace *tracer.Trace, service string) {
	if len(stack.traces) == 0 {
		stack.traces = append(stack.traces, trace)
	} else {
		stack.traces = append(stack.traces, trace)
		previous_index := len(stack.traces)-2
		previous_trace := stack.traces[previous_index]

		if previous_trace.Service == service {
			// TODO: is child! update it
			fmt.Println("hello one")
		} else {
			for _, value := range previous_trace.Attributes {
				if value == service {
					// TODO: is child! update it
					fmt.Println("hello two")
				}
			}
		}

		remove_trace(stack.traces, previous_index)
	}
}

func remove_trace(traces []*tracer.Trace, index int)  {
	fmt.Printf("%#v", traces)
	//return append(traces[:index], traces[index:]...)
}
