package main

import "fmt"

type Queue struct {
	ExpectedSequence int
	messages         map[int]Message
}

func NewQueue() Queue {
	return Queue{ExpectedSequence: 1,
		messages: make(map[int]Message)}
}

func (q *Queue) Add(msg Message) []Message {

	fmt.Println("Adding to the queue ", msg.Sequence, q.ExpectedSequence)

	q.messages[msg.Sequence] = msg

	return q.dequeue(make([]Message, 0, len(q.messages)))
}

func (q *Queue) dequeue(result []Message) []Message {

	if msg, ok := q.messages[q.ExpectedSequence]; ok {

		delete(q.messages, msg.Sequence)

		q.ExpectedSequence = msg.Sequence + 1

		return q.dequeue(append(result, msg))
	}
	return result
}
