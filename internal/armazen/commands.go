package armazen

import "oju/internal/track"

const (
	INSERT_TRACK = iota
	GET_TRACK
)

type Command interface {
	GetType() int
}

type InsertTrackCommand struct {
	command_type int
	Data         track.Track
}

type GetTrackCommand struct {
	command_type int
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

func (command GetTrackCommand) GetType() int {
	return command.command_type
}

func (command InsertTrackCommand) GetType() int {
	return command.command_type
}
