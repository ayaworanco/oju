package armazen

import "oju/internal/track"

type Action int

const (
	INSERT_TRACK Action = iota
	GET_TRACK
)

type Command interface {
	GetType() Action
}

type InsertTrackCommand struct {
	command_type Action
	Data         track.Track
}

type GetTrackCommand struct {
	command_type Action
	Data         string
}

func NewInsertTracCommand(data track.Track) InsertTrackCommand {
	return InsertTrackCommand{
		command_type: INSERT_TRACK,
		Data:         data,
	}
}

func NewGetTrackCommand(data string) GetTrackCommand {
	return GetTrackCommand{
		command_type: GET_TRACK,
		Data:         data,
	}
}

func (command GetTrackCommand) GetType() Action {
	return command.command_type
}

func (command InsertTrackCommand) GetType() Action {
	return command.command_type
}
