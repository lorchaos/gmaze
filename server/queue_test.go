package main

import (
	"testing"
)

/*
| Payload    | Sequence #| Type         | From User Id | To User Id |
|------------|-----------|--------------|--------------|------------|
|666|F|60|50 | 666       | Follow       | 60           | 50         |
|1|U|12|9    | 1         | Unfollow     | 12           | 9          |
|542532|B    | 542532    | Broadcast    | -            | -          |
|43|P|32|56  | 43        | Private Msg  | 32           | 56         |
|634|S|32    | 634       | Status Update| 32           | -          |
*/

func delivery(t *testing.T) func(m Message) {
	return func(m Message) {
		t.Logf("Got ", m.Sequence)
	}
}

func TestDelivery(t *testing.T) {

	one := Message{Payload: "", Sequence: 1, Type: "F", From: "60", To: "50"}

	queue := NewQueue()

	queue.Add(one)

	if queue.ExpectedSequence != 2 {
		t.Error("message was not delivered")
	}
}

func TestMessageOutOfOrderIsQueued(t *testing.T) {

	ten := Message{Sequence: 10}

	queue := NewQueue()

	result := queue.Add(ten)

	if len(result) != 0 {
		t.Error("nothing should be returned")
	}

	if queue.ExpectedSequence > 10 {
		t.Error("message should not be delivered")
	}
}

func TestQueuedMessageIsDelivered(t *testing.T) {

	one := Message{Payload: "", Sequence: 1, Type: "F", From: "60", To: "50"}
	two := Message{Payload: "", Sequence: 2, Type: "F", From: "60", To: "50"}

	queue := NewQueue()

	queue.Add(two)

	if queue.ExpectedSequence != 1 {
		t.Error("message should be in queue")
	}

	result := queue.Add(one)

	if l := len(result); l != 2 {
		t.Error("two messages should be in the result, got ", l)
	}

	if queue.ExpectedSequence != 3 {
		t.Error("Both messages should have been delivered")
	}
}
