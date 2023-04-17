package parser

import (
	"errors"
	"strings"
)

func Parse(packet string) (Log, error) {
	parts := strings.Split(packet, "\n")
	if len(parts) != 3 {
		return Log{}, errors.New("malformed packet")
	}

	header, header_error := NewHeader(parts[0])

	if header_error != nil {
		return Log{}, header_error
	}

	timer := parts[1]
	if timer == "" {
		return Log{}, errors.New("timer is empty")
	}
	message := parts[2]

	return Log{
		Header:  header,
		Timer:   timer,
		Message: message,
	}, nil

}
