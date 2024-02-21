package armazen

import (
	"fmt"
	"oju/internal/track"
)

const (
	INSERT_TRACK = iota
	GET_TRACK
	INVALID_COMMAND
)

type Armazen struct {
	Mailbox chan Command
}

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

func NewArmazen() Armazen {
	armazen := Armazen{
		Mailbox: make(chan Command),
	}

	go run(armazen)
	return armazen
}

func run(armazen Armazen) {
	for {
		message := <-armazen.Mailbox
		resolve(message)
	}
}

func resolve(message Command) {
	switch message.GetType() {
	case INSERT_TRACK:
		fmt.Println("inserting a track")
	case GET_TRACK:
		fmt.Println("getting a track")
	default:
		fmt.Println("doing nothing")
	}
}
