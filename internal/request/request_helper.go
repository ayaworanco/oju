package request

import (
	"fmt"
	"oju/internal/config"
)

func (header Header) String() string {
	return fmt.Sprintf("%v %v %v", header.Verb, header.Key, header.Version)
}

func (request Request) String() string {
	return fmt.Sprintf("%v\n%v\n%v", request.Header.String(), request.Timer, request.Message)
}

func is_verb_allowed(verb string) bool {
	allowed := []string{
		"LOG",
		"TRACK",
	}

	for _, allowed_verb := range allowed {
		if allowed_verb == verb {
			return true
		}
	}
	return false
}

func is_resource_allowed(key string, resources []config.Resource) bool {
	for _, resource := range resources {
		if resource.Key == key {
			return true
		}
	}
	return false
}
