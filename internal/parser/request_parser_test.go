package parser

import (
	"testing"

	"oluwoye/internal/config"
)

const TEST_PACKET = "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO\n02:49:12\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

const MALFORMED_PACKET = "LOG 5C3D47E8-2D9C-4165-A48A-7A6F6449DF66"

const MALFORMED_HEADER = "LOG AWO\n02:49:12\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

const EMPTY_TIMER = "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO\n\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

const TEST_DISALLOWED_APPS = "LOG AAAAAAAAAAAAAA AWO1.1\n02:49:12\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

const TEST_NOT_ALLOWED_VERB = "AAA AAAAAAAAAAAAAA AWO1.1\n02:49:12\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

const CONFIG_JSON = `
{
  "allowed_applications": [
    {
      "name": "worker",
      "app_key": "3FAFCF87-BF66-4DC5-84C1-34E178FF55CC"
    }
  ]
}
`

func LoadConfig() config.Config {
	config, _ := config.BuildConfig([]byte(CONFIG_JSON))
	return config
}

func TestMalformedPacket(t *testing.T) {
	config := LoadConfig()
	_, parse_error := ParseRequest(MALFORMED_PACKET, config.AllowedApplications)
	if parse_error == nil {
		t.Error("Packet not malformed")
	}
}

func TestMalformedHeader(t *testing.T) {
	config := LoadConfig()
	_, parse_error := ParseRequest(MALFORMED_HEADER, config.AllowedApplications)
	if parse_error == nil {
		t.Error("Header not malformed")
	}
}

func TestEmptyTimer(t *testing.T) {
	config := LoadConfig()
	_, parse_error := ParseRequest(EMPTY_TIMER, config.AllowedApplications)
	if parse_error == nil {
		t.Error("Timer is empty")
	}
}

func TestParseLog(t *testing.T) {
	config := LoadConfig()
	log, parse_error := ParseRequest(TEST_PACKET, config.AllowedApplications)
	if parse_error != nil {
		t.Error(parse_error.Error())
	}

	if log.Timer == "" {
		t.Error("Timer is empty")
	}
}

func TestHeaderValidationDisallowedApps(t *testing.T) {
	config := LoadConfig()
	_, parse_error := ParseRequest(TEST_DISALLOWED_APPS, config.AllowedApplications)
	if parse_error == nil {
		t.Error("Should be an error because that app is not allowed")
	}
}

func TestHeaderWithInvalidVerb(t *testing.T) {
	config := LoadConfig()
	_, parse_error := ParseRequest(TEST_NOT_ALLOWED_VERB, config.AllowedApplications)
	if parse_error == nil {
		t.Error("Should be an error because that app is not allowed")
	}
}
