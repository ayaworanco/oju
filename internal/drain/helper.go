package drain

import (
	"strings"
	"unicode"
)

func has_digit(token string) bool {
	for _, char := range token {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func is_unique_symbol(token string) bool {
	if len(token) == 1 {
		// TODO: add a regex here to check all symbols and special characters
		contains_hyphen := strings.Contains(token, "-")
		character := []rune(token)[0]
		if contains_hyphen || unicode.IsSymbol(character) {
			return true
		}
	}
	return false
}

func is_similar(sequence_1, sequence_2 []string) bool {
	n := len(sequence_1)
	if len(sequence_2) < n {
		n = len(sequence_2)
	}

	var similar_sequence float64
	for i := 0; i < n; i++ {
		if sequence_1[i] == sequence_2[i] {
			similar_sequence += 1
		}
	}
	similar_sequence = similar_sequence / float64(n)
	return similar_sequence >= SIMILARITY_THRESHHOLD
}

func get_parameters_by_similarity(sequence_1, sequence_2 []string) []string {
	var parameters []string

	n := len(sequence_1)
	if len(sequence_2) < n {
		n = len(sequence_2)
	}

	for i := 0; i < n; i++ {
		if sequence_1[i] != sequence_2[i] {
			parameters = append(parameters, sequence_1[i])
		}
	}
	return parameters
}
