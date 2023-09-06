package entities

const (
	INVALID_COMMAND = "invalid command"
	INSERT_ACTION   = "insert_action"
	GET_JOURNEY     = "get_journey"
)

type JourneyVertex struct {
	Data   string
	Target map[string]JourneyVertex
}

type Command interface {
	GetType() string
}

type InsertActionCommand struct {
	command_type string
	Data         Trace
}

type GetJourneyCommand struct {
	command_type string
	Data         string
	JourneyMap   chan map[string]JourneyVertex
}

func NewInsertActionCommand(data Trace) InsertActionCommand {
	return InsertActionCommand{
		command_type: INSERT_ACTION,
		Data:         data,
	}
}

func NewGetJourneyCommand(data string) GetJourneyCommand {
	return GetJourneyCommand{
		command_type: GET_JOURNEY,
		Data:         data,
		JourneyMap:   make(chan map[string]JourneyVertex),
	}
}

func (command GetJourneyCommand) GetType() string {
	return command.command_type
}

func (command InsertActionCommand) GetType() string {
	return command.command_type
}
