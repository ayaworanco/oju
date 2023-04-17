package parser

import (
	"errors"
	"oluwoye/internal/logger"
	"strings"
)

func Parse(packet string) (logger.Log, error) {
	parts := strings.Split(packet, "\n")
	if len(parts) > 3 || len(parts) < 3 {
		return logger.Log{}, errors.New("malformed packet")
	}

	header, header_error := logger.NewHeader(parts[0])

	if header_error != nil {
		return logger.Log{}, header_error
	}

	timer := parts[1]
	if timer == "" {
		return logger.Log{}, errors.New("timer is empty")
	}
	message := parts[2]

	return logger.Log{
		Header:  header,
		Timer:   timer,
		Message: message,
	}, nil

}
