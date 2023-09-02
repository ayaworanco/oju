package tracer

import (
	"encoding/json"
	"errors"
	"fmt"

	"oju/internal/utils"
)

const NO_TARGET_POINTED = "no target pointed"
const ESSENTIAL_DATA_EMPTY = "essential data is empty"

type Trace struct {
	id         string
	Resource   string            `json:"resource"`
	Action     string            `json:"action"`
	Target     string            `json:"target"`
	Attributes map[string]string `json:"attributes"`
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

	if tracer.Action == "" && tracer.Target == "" && tracer.Resource == "" {
		return Trace{}, errors.New(ESSENTIAL_DATA_EMPTY)
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

func (trace Trace) Print() {
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
	fmt.Println("[attributes]:")
	for key, value := range trace.Attributes {
		fmt.Printf("\t[%s]: %s\n", key, value)
	}
}
