package request

import (
	"errors"
	"oju/internal/config"
	"strings"
	"time"
)

const (
	ERROR_RESOURCE_NOT_ALLOWED = "resource_not_allowed"
	ERROR_MALFORMED_HEADER     = "malformed_header"
	ERROR_MALFORMED_PACKET     = "malformed_packet"
	ERROR_VERB_NOT_ALLOWED     = "verb_not_allowed"
)

type Request struct {
	Header         Header
	Timer, Message string
}

type Header struct {
	Verb, Key, Version string
}

// This function builds a request based on a raw TCP packet
func new_request(head, message string, resources []config.Resource) (Request, error) {
	parts := strings.Split(head, " ")
	if len(parts) != 3 {
		return Request{}, errors.New(ERROR_MALFORMED_HEADER)
	}

	verb := parts[0]
	if !is_verb_allowed(verb) {
		return Request{}, errors.New(ERROR_VERB_NOT_ALLOWED)
	}

	key := parts[1]
	if !is_resource_allowed(key, resources) {
		return Request{}, errors.New(ERROR_RESOURCE_NOT_ALLOWED)
	}

	version := parts[2]

	header := Header{Verb: verb, Key: key, Version: version}

	return Request{
		Header:  header,
		Timer:   time.Now().String(),
		Message: message,
	}, nil
}

func Parse(packet string, resources []config.Resource) (Request, error) {
	// TODO: try to understand if is a HTTP, Raw TCP or gRPC request
	parts := strings.Split(packet, "\n")
	if len(parts) != 2 {
		return Request{}, errors.New(ERROR_MALFORMED_PACKET)
	}

	head := parts[0]
	message := parts[1]

	request, request_error := new_request(head, message, resources)
	if request_error != nil {
		return Request{}, request_error
	}

	return request, nil
}
