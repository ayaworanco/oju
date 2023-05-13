package querier

import (
	"regexp"

	"oluwoye/internal/parser"
)

type node struct {
	Data   map[string]interface{}
	Result bool
}

func (node *node) set_result(log_groups []*parser.LogGroup) {
	regex := node.Data["variable"]
	operator := node.Data["operator"]
	value := node.Data["value"]

	var variable_result string

	if node.is_term() {
		expression := regex.(regexp.Regexp)
		for _, group := range log_groups {
			for _, parameters := range group.LogParameters {
				for _, parameter := range parameters {
					if expression.MatchString(parameter) {
						variable_result = parameter
						break
					}
				}
			}
		}
		node.Result = logical_operation(variable_result, value, operator.(string))
	} else {
		node.Result = false
	}
}

func (node *node) is_term() bool {
	_, ok_variable := node.Data["variable"]
	_, ok_operator := node.Data["operator"]
	_, ok_value := node.Data["value"]
	return ok_variable && ok_operator && ok_value
}

func (node *node) is_operator() bool {
	_, ok_variable := node.Data["variable"]
	_, ok_operator := node.Data["operator"]
	_, ok_value := node.Data["value"]
	return !ok_variable && ok_operator && !ok_value
}
