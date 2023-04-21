package ruler

import (
	"errors"
	"regexp"

	"gopkg.in/yaml.v3"
)

var RESOURCE_MAP = map[string]regexp.Regexp{
	"$ipv4": *regexp.MustCompile("^(?P<ipv4>[0-9]+.[0-9]+.[0-9]+.[0-9]+)"),
}

type Rule struct {
	Resource string `yaml:"resource"`
	Target   string `yaml:"target"`
	Operator string `yaml:"operator"`
	Action   Action `yaml:"action"`
}

func (rule *Rule) Run(message string) (bool, error) {
	resource, build_resource_error := build_resource(rule.Resource, message)
	if build_resource_error != nil {
		return false, build_resource_error
	}

	target, build_target_error := build_target(rule.Target)
	if build_target_error != nil {
		return false, build_target_error
	}

	switch rule.Operator {
	case "equal":
		return equal(resource, target), nil
	default:
		return false, errors.New("unknown operator while running rule")
	}
}

func LoadRules(raw []byte) ([]Rule, error) {
	var rules []Rule
	err := yaml.Unmarshal(raw, &rules)
	if err != nil {
		return []Rule{}, err
	}

	return rules, nil
}
