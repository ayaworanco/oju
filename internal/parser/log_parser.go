package parser

import (
	"errors"
	"oluwoye/internal/config"
	"strings"
)

func Parse(packet string, allowed_applications []config.Application) (Log, error) {
	parts := strings.Split(packet, "\n")
	if len(parts) != 3 {
		return Log{}, errors.New("malformed packet")
	}

	header, header_error := NewHeader(parts[0], allowed_applications)

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
