package journey

import (
	"oju/internal/domain/entities"
	"testing"
)

func TestAddJourneyDataInEmptyJourney(t *testing.T) {
	test_journey := entities.Journey{
		Children: make(map[string][]entities.JourneyData),
	}

	new_journey := AddData(test_journey, "some_action", entities.JourneyData{
		Action:     "some_action",
		Service:    "some_service",
		Attributes: make(map[string]string),
	})

	if len(new_journey.Children) == 0 {
		t.Error("journey should be filled")
	}
}
