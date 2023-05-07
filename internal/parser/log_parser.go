package parser

import (
	"strings"
	"unicode"
)

const SIMILARITY_THRESHHOLD = 0.6
const MAX_CHILD = 100

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

func has_digit(token string) bool {
	for _, char := range token {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func get_parameter_by_similarity(sequence_1, sequence_2 []string) string {
	n := len(sequence_1)
	if len(sequence_2) < n {
		n = len(sequence_2)
	}

	for i := 0; i < n; i++ {
		if sequence_1[i] != sequence_2[i] {
			return sequence_1[i]
		}
	}
	return ""
}

func is_similar(sequence_1, sequence_2 []string) bool {
	n := len(sequence_1)
	if len(sequence_2) < n {
		n = len(sequence_2)
	}

	var simSeq float64
	for i := 0; i < n; i++ {
		if sequence_1[i] == sequence_2[i] {
			simSeq += 1
		}
	}
	simSeq = simSeq / float64(n)
	return simSeq >= SIMILARITY_THRESHHOLD
}

func ParseLog(tree *Tree, log string, id int) {
	tree.add_or_update_length_layer(log, id)
}
