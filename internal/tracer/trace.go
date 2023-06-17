package tracer

import (
	"encoding/json"

	"oju/internal/utils"
)

type Trace struct {
	id         string
	AppKey     string            `json:"app_key"`
	Name       string            `json:"name"`
	Service    string            `json:"service"`
	Attributes map[string]string `json:"attributes"`
	children   map[string]*Trace
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

func (trace *Trace) GetId() string {
	return trace.id
}

func (trace *Trace) GetChildren() map[string]*Trace {
	return trace.children
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
