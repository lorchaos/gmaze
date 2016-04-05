package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Target map[string]bool

type MessageDelivery struct {
	Target  Target
	Payload string
}

type State struct {
	followers map[string]Target
	all       Target
}

func NewState() State {
	return State{make(map[string]Target),
		make(Target),
	}
}

func (s State) Register(id string) {

}

/*
Follow: Only the To User Id should be notified
Unfollow: No clients should be notified
Broadcast: All connected user clients should be notified
Private Message: Only the To User Id should be notified
Status Update: All current followers of the From User ID should be notified
*/
func (s State) Process(msg Message) MessageDelivery {

	target := s.process(msg)

	return MessageDelivery{target, msg.Payload}
}

func (s State) process(msg Message) Target {

	s.all[msg.To] = true
	s.all[msg.From] = true

	if msg.Type == "F" {

		m := s.followers[msg.To]
		if m == nil {
			m = make(Target)
			s.followers[msg.To] = m
		}
		m[msg.From] = true

		return map[string]bool{msg.To: true}

	} else if msg.Type == "U" {

		m := s.followers[msg.To]
		if m == nil {
			m = make(Target)
			s.followers[msg.To] = m
		}

		m[msg.From] = false

	} else if msg.Type == "B" {

		return nil

	} else if msg.Type == "S" {

		m := s.followers[msg.From]

		if m == nil {
			return make(Target)
		}
		return m

	} else if msg.Type == "P" {

		return map[string]bool{msg.To: true}
	}

	return map[string]bool{}
}

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

type Delivery func(Message)

type Queue struct {
	ExpectedSequence int
	delivery         Delivery
	messages         map[int]Message
}

type MessageQueue []Message

func (a MessageQueue) Len() int           { return len(a) }
func (a MessageQueue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MessageQueue) Less(i, j int) bool { return a[i].Sequence < a[j].Sequence }

func NewQueue() Queue {
	return Queue{ExpectedSequence: 1,
		messages: make(map[int]Message)}
}

func (q *Queue) Add(msg Message) []Message {

	if len(q.messages) > 0 {
		fmt.Printf("Msg %d exp %d size %d tip %d\n", msg.Sequence, q.ExpectedSequence, len(q.messages), q.messages[0].Sequence)
	}

	fmt.Println("Adding to the queue ", msg.Sequence, q.ExpectedSequence)

	q.messages[msg.Sequence] = msg

	result := q.verifyQueue()

	if len(result) > 0 {
		fmt.Println("Dequeued ", len(result))
	}
	return result
}

func (q *Queue) verifyQueue() []Message {

	result := make([]Message, 0, len(q.messages))

	for {

		if msg, ok := q.messages[q.ExpectedSequence]; ok {

			//fmt.Println("Message found ", q.ExpectedSequence)

			result = append(result, msg)

			delete(q.messages, msg.Sequence)

			q.ExpectedSequence = msg.Sequence + 1

		} else {
			return result
		}
	}

	return result
}
