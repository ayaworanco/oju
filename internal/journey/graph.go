package journey

import (
	"oju/internal/config"
	"oju/internal/tracer"
	"oju/internal/utils"
	"fmt"
)

type graph struct {
	vertices     map[string]vertex
	applications []config.Application // INFO: this is needed?
}

type vertex struct {
	name string
}

func new_graph(applications []config.Application) graph {
	return graph{
		vertices:     make(map[string]vertex),
		applications: applications,
	}
}

func new_graph_pure(vertices map[string]vertex) graph {
	return graph{
		vertices: vertices,
		applications: []config.Application{},
	}
}

func update_graph(old_graph graph, data tracer.Trace) graph {
	new_vertex := vertex{name: data.Name}

	action := fmt.Sprintf("%s@%s", data.AppKey, data.Name)
	new_map := utils.MapPut(old_graph.vertices, action, new_vertex)

	return new_graph_pure(new_map)
}
