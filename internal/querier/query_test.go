package querier

import "testing"

const LOG_MESSAGE = "54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

func TestValidQuery(t *testing.T) {

	input := "'$ipv4 eq 54.36.149.41 and $status_code eq 200'"

	result, error := Parse(input, LOG_MESSAGE)
	if error != nil {
		t.Error(error.Error())
	}

	if !result {
		t.Error("Scenario not valid")
	}
}
