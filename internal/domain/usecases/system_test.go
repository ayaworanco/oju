package usecases

import (
	"oju/internal/domain/entities"
	"testing"
)

type test_suite struct {
	resources []entities.Resource
	system    entities.System
}

func setup() test_suite {
	resources := []entities.Resource{
		{
			Name: "test_a",
			Key:  "test_a",
			Host: "http://test_a.svc.cluster.local",
		},
	}

	return test_suite{
		resources: resources,
		system:    NewSystem(resources),
	}
}

func TestInsertAction(t *testing.T) {
	suite := setup()

	trace := entities.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "delta",
		Attributes: make(map[string]string),
	}

	message := entities.NewInsertActionCommand(trace)

	if message.Data.Action == "" {
		t.Error("trace should be fulfilled")
	}

	Send(suite.system, message)
}

func TestGetJourney(t *testing.T) {
	suite := setup()

	trace := entities.Trace{
		Resource:   "test1",
		Action:     "bhaskara",
		Target:     "test2@delta",
		Attributes: make(map[string]string),
	}
	trace_delta := entities.Trace{
		Resource:   "test2",
		Action:     "delta",
		Target:     "",
		Attributes: make(map[string]string),
	}

	message := entities.NewInsertActionCommand(trace)
	message_delta := entities.NewInsertActionCommand(trace_delta)
	message_get_journey_bhaskara := entities.NewGetJourneyCommand("test1@bhaskara")

	Send(suite.system, message)
	Send(suite.system, message_delta)
	Send(suite.system, message_get_journey_bhaskara)

	journey_map := <-message_get_journey_bhaskara.JourneyMap

	close(message_get_journey_bhaskara.JourneyMap)

	if len(journey_map) == 0 {
		t.Error("Journey cannot be empty")
	}
}
