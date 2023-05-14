package parser

import (
	"errors"
	"strings"

	"oluwoye/internal/config"
)

func ParseRequest(packet string, allowed_applications []config.Application) (Request, error) {
	parts := strings.Split(packet, "\n")
	if len(parts) != 3 {
		return Request{}, errors.New("malformed packet")
	}

	header, header_error := NewHeader(parts[0], allowed_applications)

	if header_error != nil {
		return Request{}, header_error
	}

	timer := parts[1]
	if timer == "" {
		return Request{}, errors.New("timer is empty")
	}
	message := parts[2]

	return Request{
		Header:  header,
		Timer:   timer,
		Message: message,
	}, nil
}
