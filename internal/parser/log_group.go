package parser

import (
	"fmt"
	"strings"
)

// TODO: generate a unique identifier md5 based on time

type LogGroup struct {
	LogEvent      string
	LogParameters map[string][]string
}

func new_log_group(log_event string, log_parameters map[string][]string) *LogGroup {
	return &LogGroup{
		LogEvent:      log_event,
		LogParameters: log_parameters,
	}
}

func add_log_group(node *node, log_message string, id string) {
	sequence_log_message := strings.Split(log_message, " ")

	log_event := log_message
	var first_parameters []string

	for index, token := range sequence_log_message {
		if has_digit(token) {
			first_parameters = append(first_parameters, token)
			sequence_log_message[index] = "*"
			log_event = strings.Join(sequence_log_message, " ")
		}
	}

	var log_group *LogGroup

	if len(first_parameters) > 0 {
		log_group = new_log_group(log_event, map[string][]string{id: first_parameters})
	} else {
		log_group = new_log_group(log_event, map[string][]string{})
	}

	log_group_id := fmt.Sprintf("log_group_%v", len(strings.Split(log_message, " ")))

	child, ok := node.children[log_group_id]
	if !ok {
		child = new_node(log_group_id, log_group)
		node.children[log_group_id] = child
	}

	found_log_group := child.data.(*LogGroup)
	sequence_1 := strings.Split(log_message, " ")
	sequence_2 := strings.Split(found_log_group.LogEvent, " ")

	if is_similar(sequence_1, sequence_2) {
		parameters := get_parameters_by_similarity(sequence_1, sequence_2)
		if len(parameters) > 0 {
			found_log_group.LogParameters[id] = append(found_log_group.LogParameters[id], parameters...)
		}
		update_log_event(found_log_group, log_message)
	}
}
