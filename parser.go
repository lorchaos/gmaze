package gmaze

import (
	"strings"
	"strconv"
	)

type Message struct {
	Payload string
	Sequence int
	Type string
	From string
	To string
} 

func Parse(msg string) Message {

	fields := strings.Split(msg, "|")
	msgType := fields[1]

	seq, _ := strconv.Atoi(fields[0])

	result := Message { Payload: msg, Sequence: seq, Type: msgType }

	if len(fields) >= 3 {
		result.From =  fields[2]
	}

	if len(fields) >= 4 {
		result.To = fields[3]
	}

	return result
}

type Delivery func(Message)

type Queue struct {
	ExpectedSequence int
	delivery Delivery
}

func NewQueue(delivery Delivery) Queue {
	return Queue { ExpectedSequence: 0,  delivery: delivery}
}

func (q* Queue) Add(msg Message) {

	if msg.Sequence == q.ExpectedSequence {
		
		q.ExpectedSequence = q.ExpectedSequence +1
		q.delivery(msg)
	}

}