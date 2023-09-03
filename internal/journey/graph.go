package journey

import (
	"oju/internal/utils"

	"fmt"
)

type Graph struct {
	Vertices map[string]Vertex
}

type Vertex struct {
	Name   string
	Target string
}

func NewGraph(vertices map[string]Vertex) Graph {
	return Graph{
		Vertices: vertices,
	}
}

func UpdateGraph(old_graph Graph, command InsertActionCommand) Graph {
	action := fmt.Sprintf("%s@%s", command.Data.Resource, command.Data.Action)

	new_vertex := Vertex{Name: command.Data.Action, Target: command.Data.Target}
	new_map := utils.MapPut(old_graph.Vertices, action, new_vertex)

	return NewGraph(new_map)
}

func GetJourney(action string, current map[string]Vertex, vertices map[string]Vertex) map[string]journey_vertex {
	new_journey := map[string]journey_vertex{action: {data: action}}

	if current[action].Target == "" {
		return new_journey
	}

	target_name := current[action].Target
	target_vertex := vertices[target_name]

	if entry, ok := new_journey[action]; ok {
		entry.target = GetJourney(target_name, map[string]Vertex{target_name: target_vertex}, vertices)
		new_journey[action] = entry
	}

	return new_journey
}
