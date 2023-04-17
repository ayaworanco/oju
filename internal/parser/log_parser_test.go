package parser

import (
	"testing"
)

const TEST_PACKET = "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO\n02:49:12\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

const MALFORMED_PACKET = "LOG 5C3D47E8-2D9C-4165-A48A-7A6F6449DF66"

const MALFORMED_HEADER = "LOG AWO\n02:49:12\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

const EMPTY_TIMER = "LOG 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO\n\n54.36.149.41 - - [22/Jan/2019:03:56:14 +0330] \"GET /filter/27|13%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,27|%DA%A9%D9%85%D8%AA%D8%B1%20%D8%A7%D8%B2%205%20%D9%85%DA%AF%D8%A7%D9%BE%DB%8C%DA%A9%D8%B3%D9%84,p53 HTTP/1.1\" 200 30577 \"-\" \"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)\" \"-\""

func TestMalformedPacket(t *testing.T) {
	_, parse_error := Parse(MALFORMED_PACKET)
	if parse_error == nil {
		t.Error("Packet not malformed")
	}
}

func TestMalformedHeader(t *testing.T) {
	_, parse_error := Parse(MALFORMED_HEADER)
	if parse_error == nil {
		t.Error("Header not malformed")
	}
}

func TestEmptyTimer(t *testing.T) {
	_, parse_error := Parse(EMPTY_TIMER)
	if parse_error == nil {
		t.Error("Timer is empty")
	}
}

func TestParseLog(t *testing.T) {
	log, parse_error := Parse(TEST_PACKET)
	if parse_error != nil {
		t.Error(parse_error.Error())
	}

	if log.Timer == "" {
		t.Error("Timer is empty")
	}
}
