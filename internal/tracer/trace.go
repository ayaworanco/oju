package tracer

import (
	"encoding/json"
	"fmt"

	"oju/internal/utils"
)

const NO_TARGET_POINTED = "no target pointed"

type Trace struct {
	id         string
	Resource   string            `json:"resource"`
	Action     string            `json:"action"`
	Target     string            `json:"target"`
	Attributes map[string]string `json:"attributes"`
	children   map[string]*Trace
}

type IsTrace interface {
	Trace
	GetId() string
}

func Parse(packet string) (Trace, error) {
	var tracer Trace
	unmarshal_error := json.Unmarshal([]byte(packet), &tracer)

	if unmarshal_error != nil {
		return Trace{}, unmarshal_error
	}

	tracer.SetId()

	return tracer, nil
}

func (trace *Trace) SetId() {
	trace.id = utils.GenerateId()
}

func (trace Trace) GetId() string {
	return trace.id
}

func (trace *Trace) GetChildren() map[string]*Trace {
	return trace.children
}

func (trace *Trace) SetResource(resource string) {
	trace.Resource = resource
}

func (trace *Trace) AddChild(new_trace *Trace) {
	id := new_trace.GetId()
	if trace.children == nil {
		trace.children = make(map[string]*Trace)
		trace.children[id] = new_trace
	} else {
		trace.children[id] = new_trace
	}
}

func (trace *Trace) Print() {
	var service string

	if trace.Target == "" {
		service = NO_TARGET_POINTED
	} else {
		service = trace.Target
	}

	fmt.Println("=> TRACE from ", trace.Resource)
	fmt.Println("[id]: ", trace.GetId())
	fmt.Println("[action]: ", trace.Action)

	fmt.Println("[service]: ", service)
	fmt.Println("[children]: ", len(trace.GetChildren()))
	fmt.Println("[attributes]:")
	for key, value := range trace.Attributes {
		fmt.Printf("\t[%s]: %s\n", key, value)
	}
}
