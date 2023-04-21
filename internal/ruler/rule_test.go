package ruler

import (
	"testing"
)

const TEST_PACKET = "54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

func TestRunRule(t *testing.T) {
	rule := Rule{
		Resource: "$ipv4",
		Target:   "54.36.149.41",
		Operator: "equal",
		Action: Action{
			Name: "action_by_mail",
			Parameters: []string{
				"tech_lead@gmail.com",
				"warning",
			},
		},
	}

	result, rule_error := rule.Run(TEST_PACKET)
	if rule_error != nil {
		t.Error(rule_error.Error())
	}

	if !result {
		t.Error("Rule must be true for alerting")
	}
}

func TestRunRuleAndReturnFalse(t *testing.T) {
	rule := Rule{
		Resource: "$ipv4",
		Target:   "54.36.149.42",
		Operator: "equal",
		Action: Action{
			Name: "action_by_mail",
			Parameters: []string{
				"tech_lead@gmail.com",
				"warning",
			},
		},
	}

	result, rule_error := rule.Run(TEST_PACKET)
	if rule_error != nil {
		t.Error(rule_error.Error())
	}

	if result {
		t.Error("Rule must be false")
	}
}

func TestLoadRulesAndRunInPacket(t *testing.T) {
	rules_yaml := `
- resource: $ipv4
  operator: equal
  target: 54.36.149.41
  action:
    name: alert_by_email
    parameters:
      - tech_lead@gmail.com
      - warning
`
	rules, load_error := LoadRules([]byte(rules_yaml))
	if load_error != nil {
		t.Error("Rules not loaded")
	}

	if len(rules) == 0 {
		t.Error("Rules is loaded and is empty")
	}
}
