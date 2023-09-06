package usecases

import (
	"errors"
	"oju/internal/domain/entities"
	"strings"
	"time"
)

const (
	ERROR_APPLICATION_NOT_ALLOWED = "application not allowed"
	ERROR_MALFORMED_HEADER        = "malformed_header"
	ERROR_MALFORMED_PACKET        = "malformed packet"
	ERROR_VERB_NOT_ALLOWED        = "verb_not_allowed"
)

func NewRequest(head, message string, resources []entities.Resource) (entities.Request, error) {
	parts := strings.Split(head, " ")
	if len(parts) != 3 {
		return entities.Request{}, errors.New(ERROR_MALFORMED_HEADER)
	}

	verb := parts[0]
	if !is_verb_allowed(verb) {
		return entities.Request{}, errors.New(ERROR_VERB_NOT_ALLOWED)
	}

	app_key := parts[1]

	if !is_application_allowed(app_key, resources) {
		return entities.Request{}, errors.New(ERROR_APPLICATION_NOT_ALLOWED)
	}

	version := parts[2]

	header := entities.Header{Verb: verb, AppKey: app_key, Version: version}

	return entities.Request{
		Header:  header,
		Timer:   time.Now().String(),
		Message: message,
	}, nil
}

func Parse(packet string, resources []entities.Resource) (entities.Request, error) {
	parts := strings.Split(packet, "\n")
	if len(parts) != 2 {
		return entities.Request{}, errors.New(ERROR_MALFORMED_PACKET)
	}

	head := parts[0]
	message := parts[1]

	request, request_error := NewRequest(head, message, resources)
	if request_error != nil {
		return entities.Request{}, request_error
	}
	return request, nil
}

func is_verb_allowed(verb string) bool {
	allowed := []string{
		"LOG",
		"WATCH",
		"TRACE",
	}

	for _, allowed_verb := range allowed {
		if allowed_verb == verb {
			return true
		}
	}
	return false
}

func is_application_allowed(app_key string, resources []entities.Resource) bool {

	for _, application := range resources {
		if application.Key == app_key {
			return true
		}
	}
	return false
}
