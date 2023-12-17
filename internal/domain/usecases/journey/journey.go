package journey

import (
	"maps"
	"oju/internal/domain/entities"
)

func AddData(journey entities.Journey, resource, action string, data entities.JourneyData) entities.Journey {
	resources := journey.Resources
	cloned := maps.Clone(resources[resource])
	children := maps.Clone(journey.Applications)
	children[action] = append(children[action], data)
	new_journey := entities.Journey{Children: children}

	return new_journey
}
