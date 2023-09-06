package entities

import "fmt"

type Request struct {
	Header         Header
	Timer, Message string
}

type Header struct {
	Verb, AppKey, Version string
}

func (header Header) String() string {
	return fmt.Sprintf("%v %v %v", header.Verb, header.AppKey, header.Version)
}

func (request Request) String() string {
	return fmt.Sprintf("%v\n%v\n%v", request.Header.String(), request.Timer, request.Message)
}
