package usecases

import (
	"fmt"
	"oju/internal/domain/entities"
	"oju/internal/utils"
)

func new_graph(vertices map[string]entities.Vertex) entities.Graph {
	return entities.Graph{
		Vertices: vertices,
	}
}

func update_graph(old_graph entities.Graph, command entities.InsertActionCommand) entities.Graph {
	action := fmt.Sprintf("%s@%s", command.Data.Resource, command.Data.Action)

	new_vertex := entities.Vertex{Name: command.Data.Action, Target: command.Data.Target}
	new_map := utils.MapPut[string, entities.Vertex](old_graph.Vertices, action, new_vertex)

	return new_graph(new_map)
}

func get_journey(action string, current map[string]entities.Vertex, vertices map[string]entities.Vertex) map[string]entities.JourneyVertex {
	new_journey := map[string]entities.JourneyVertex{action: {Data: action}}

	if current[action].Target == "" {
		return new_journey
	}

	target_name := current[action].Target
	target_vertex := vertices[target_name]

	if entry, ok := new_journey[action]; ok {
		entry.Target = get_journey(target_name, map[string]entities.Vertex{target_name: target_vertex}, vertices)
		new_journey[action] = entry
	}

	return new_journey
}
