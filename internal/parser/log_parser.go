package parser

import (
	"fmt"
	"strings"
	"unicode"
)

const SIMILARITY_THRESHHOLD = 0.6
const MAX_CHILD = 100

type Node struct {
	Data     interface{}
	Label    string
	Children map[string]*Node
}

type Tree struct {
	Root  *Node
	Depth int
}

type LogGroup struct {
	LogEvent      string
	LogParameters map[int]string
}

func NewTree(depth int) *Tree {
	return &Tree{
		Root: &Node{
			Data:     "root",
			Children: make(map[string]*Node, 0),
		},
		Depth: depth - 2,
	}
}

func NewNode(label string, data interface{}) *Node {
	return &Node{
		Data:     data,
		Label:    label,
		Children: make(map[string]*Node),
	}
}

func new_log_group(log_event string, log_parameter map[int]string) *LogGroup {
	return &LogGroup{
		LogEvent:      log_event,
		LogParameters: log_parameter,
	}
}

func (tree *Tree) AddOrUpdateLengthLayer(log string, id int) {
	parts := strings.SplitN(log, " ", tree.Depth)
	length := len(parts)
	label := fmt.Sprint(length)

	child, ok := tree.Root.Children[label]
	if !ok {
		child = NewNode(label, length)
		tree.Root.Children[label] = child

	}
	child.Add(parts, log, id)
}

func (node *Node) Add(parts []string, log_message string, id int) {
	if len(node.Children) > MAX_CHILD {
		return
	}

	if len(parts) == 0 {
		add_log_group(node, log_message, id)
		return
	}

	var path string
	if len(parts) == 1 {
		path = parts[0]
	} else {
		path = parts[:len(parts)-1][0]
	}

	if has_digit(path) {
		path = "*"
	}

	child, ok := node.Children[path]
	if !ok {
		child = NewNode(path, path)
		node.Children[path] = child
	}

	child.Add(parts[1:], log_message, id)
}

func add_log_group(node *Node, log_message string, id int) {
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

	child, ok := node.Children[log_group_id]
	if !ok {
		child = NewNode(log_group_id, log_group)
		node.Children[log_group_id] = child
	}

	found_log_group := child.Data.(*LogGroup)
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
	tree.AddOrUpdateLengthLayer(log, id)
}
