package ruler

// Variables
type Ipv4 string
type Date string
type Verb string
type StatusCode int
type Path string
type Unknown string

type VariableType interface {
	comparable
	Ipv4 | Date | Verb | StatusCode | Path | Unknown
}

type Variable struct {
	Value interface{}
}

type Action struct {
	Name       string
	Parameters []string
}
