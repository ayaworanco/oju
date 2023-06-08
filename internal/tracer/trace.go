package tracer

import (
	"encoding/json"

	"oju/internal/utils"
)

type Trace struct {
	id         string
	Name       string      `json:"name"`
	Service    string      `json:"service"`
	Attributes interface{} `json:"attributes"`
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
