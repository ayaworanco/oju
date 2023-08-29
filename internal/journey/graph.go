package journey

import (
	"oju/internal/config"
	"oju/internal/tracer"
	//"oju/internal/utils"
	"fmt"
)

const (
	INVALID_COMMAND = "invalid command"
)

type graph struct {
	vertices     map[string]vertex
	applications []config.Application // INFO: this is needed?
}

type vertex struct {
	name string
}

type InsertActionCommand struct {
	Type string
	Data tracer.Trace
}

func new_graph(vertices map[string]vertex) graph {
	return graph{
		vertices: vertices,
		applications: []config.Application{},
	}
}

func update_graph(old_graph graph, command InsertActionCommand) graph {
	fmt.Printf("%#v", command)
	/*
	new_vertex := vertex{name: command.Data.Name}

	action := fmt.Sprintf("%s@%s", command.Data.AppKey, command.Data.Name)
	new_map := utils.MapPut(old_graph.vertices, action, new_vertex)

	return new_graph(new_map)
	*/
	return old_graph
}
