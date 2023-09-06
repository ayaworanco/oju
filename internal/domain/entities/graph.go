package entities

type Graph struct {
	Vertices map[string]Vertex
}

type Vertex struct {
	Name, Target string
}
