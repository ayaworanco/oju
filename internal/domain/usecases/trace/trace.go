package usecases

import (
	"encoding/json"
	"errors"
	"oju/internal/domain/entities"
)

const (
	NO_TARGET_POINTED    = "no target pointed"
	ESSENTIAL_DATA_EMPTY = "essential data is empty"
)

func ParseTrace(packet string) (entities.Trace, error) {
	var tracer entities.Trace
	unmarshal_error := json.Unmarshal([]byte(packet), &tracer)

	if unmarshal_error != nil {
		return entities.Trace{}, unmarshal_error
	}

	if tracer.Action == "" && tracer.Target == "" && tracer.Resource == "" {
		return entities.Trace{}, errors.New(ESSENTIAL_DATA_EMPTY)
	}

	tracer.SetId()

	return tracer, nil
}
