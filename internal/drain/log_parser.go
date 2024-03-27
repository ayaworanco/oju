package drain

import (
	"strings"
)

const (
	SIMILARITY_THRESHHOLD = 0.6
	MAX_CHILD             = 100
)

func update_log_event(log_group *LogGroup, log_message string) {
	sequence_log_event := strings.Split(log_group.LogEvent, " ")
	sequence_log_message := strings.Split(log_message, " ")

	n := len(sequence_log_event)
	if len(sequence_log_message) < n {
		n = len(sequence_log_message)
	}

	for i := 0; i < n; i++ {
		if sequence_log_event[i] != sequence_log_message[i] {
			sequence_log_event[i] = "*"
			log_group.LogEvent = strings.Join(sequence_log_event, " ")
			break
		}
	}
}

func remount_without_symbols(tokens []string) []string {
	var new_tokens []string
	for _, token := range tokens {
		if !is_unique_symbol(token) {
			new_tokens = append(new_tokens, token)
		}
	}
	return new_tokens
}

func ParseLog(tree *Tree, log string) {
	tree.add_or_update_length_layer(log)
}
