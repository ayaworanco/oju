package ruler

import (
	"errors"
	"regexp"
)

func equal(resource interface{}, target interface{}) bool {
	return resource == target
}

func build_target(target string) (Variable, error) {
	var resources []Variable

	for key, regex := range RESOURCE_MAP {
		variable, error_variable := make_variable(key, regex, target)
		if error_variable != nil {
			continue
		} else {
			resources = append(resources, variable)
			break
		}
	}
	return resources[0], nil
}

func build_resource(resource string, message string) (Variable, error) {
	var resource_key string
	var regexp regexp.Regexp

	for key, value := range RESOURCE_MAP {
		if resource == key {
			resource_key = key
			regexp = value
			break
		}
	}

	variable, make_variable_error := make_variable(resource_key, regexp, message)

	if make_variable_error != nil {
		return Variable{}, make_variable_error
	}
	return variable, nil

}

func make_variable(key string, regex regexp.Regexp, value string) (Variable, error) {
	switch key {
	case "$ipv4":
		if regex.MatchString(value) {
			found := regex.FindString(value)
			return Variable{
				Value: Ipv4(found),
			}, nil
		} else {
			return Variable{Value: Unknown("")}, errors.New("unknown variable")
		}
	default:
		return Variable{Value: Unknown("")}, errors.New("unknown variable")
	}
}
