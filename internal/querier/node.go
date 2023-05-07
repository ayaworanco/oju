package querier

import "regexp"

type node struct {
	Data   map[string]interface{}
	Result bool
}

func (node *node) set_result(message string) {
	regex := node.Data["variable"]
	operator := node.Data["operator"]
	value := node.Data["value"]

	var variable_result string

	if node.is_term() {
		expression := regex.(regexp.Regexp)
		if expression.MatchString(message) {
			variable_result = expression.FindString(message)
		} else {
			node.Result = false
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
