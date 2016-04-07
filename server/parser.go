package main

import (
	"strconv"
	"strings"
)

type Message struct {
	Payload  string
	Sequence int
	Type     string
	From     string
	To       string
}

func Parse(msg string) Message {

	fields := strings.Split(msg, "|")
	msgType := fields[1]

	seq, _ := strconv.Atoi(fields[0])

	result := Message{Payload: msg, Sequence: seq, Type: msgType}

	if len(fields) >= 3 {
		result.From = fields[2]
	}

	if len(fields) >= 4 {
		result.To = fields[3]
	}

	return result
}
