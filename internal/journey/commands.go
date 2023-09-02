package journey

import "oju/internal/tracer"

const (
	INVALID_COMMAND = "invalid command"
)

const (
	INSERT_ACTION = "insert_action"
	GET_JOURNEY   = "get_journey"
)

type journey_vertex struct {
	data   string
	target map[string]journey_vertex
}

type Command interface {
	GetType() string
}

type InsertActionCommand struct {
	command_type string
	Data         tracer.Trace
}

type GetJourneyCommand struct {
	command_type string
	Data         string
	JourneyMap   chan map[string]journey_vertex
}

func NewInsertActionCommand(data tracer.Trace) InsertActionCommand {
	return InsertActionCommand{
		command_type: INSERT_ACTION,
		Data:         data,
	}
}

func NewGetJourneyCommand(data string) GetJourneyCommand {
	return GetJourneyCommand{
		command_type: GET_JOURNEY,
		Data:         data,
		JourneyMap:   make(chan map[string]journey_vertex),
	}
}

func (command GetJourneyCommand) GetType() string {
	return command.command_type
}

func (command InsertActionCommand) GetType() string {
	return command.command_type
}
