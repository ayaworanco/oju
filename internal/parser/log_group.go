package parser

import (
	"fmt"
	"strings"
)

type LogGroup struct {
	LogEvent      string
	LogParameters map[int]string
}

func new_log_group(log_event string, log_parameter map[int]string) *LogGroup {
	return &LogGroup{
		LogEvent:      log_event,
		LogParameters: log_parameter,
	}
}

func add_log_group(node *node, log_message string, id int) {
	sequence_log_message := strings.Split(log_message, " ")

	log_event := log_message
	first_parameter := ""

	for i, token := range sequence_log_message {
		if has_digit(token) {
			first_parameter = token
			sequence_log_message[i] = "*"
			log_event = strings.Join(sequence_log_message, " ")
			break
		}
	}

	var log_group *LogGroup

	if first_parameter != "" {
		log_group = new_log_group(log_event, map[int]string{id: first_parameter})
	} else {
		log_group = new_log_group(log_event, map[int]string{})
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
		parameter := get_parameter_by_similarity(sequence_1, sequence_2)
		if parameter != "" {
			found_log_group.LogParameters[id] = parameter
		}
		update_log_event(found_log_group, log_message)
	}
}
