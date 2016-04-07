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

func TestFollowParser(t *testing.T) {

	payload := "666|F|60|50"
	msg := Parse(payload)

	expected := Message{Payload: payload, Sequence: 666, Type: "F", From: "60", To: "50"}

	if msg != expected {
		t.Fail()
	}
}

func TestUnfollowParser(t *testing.T) {

	payload := "1|U|12|9"
	msg := Parse(payload)

	expected := Message{Payload: payload, Sequence: 1, Type: "U", From: "12", To: "9"}

	if msg != expected {
		t.Fail()
	}
}

func TestBroadcastParser(t *testing.T) {

	payload := "542532|B"
	msg := Parse(payload)

	expected := Message{Payload: payload, Sequence: 542532, Type: "B"}

	if msg != expected {
		t.Fail()
	}
}

func TestPrivateParser(t *testing.T) {

	payload := "43|P|32|56"
	msg := Parse(payload)

	expected := Message{Payload: payload, Sequence: 43, Type: "P", From: "32", To: "56"}

	if msg != expected {
		t.Fail()
	}
}

func TestStatusUpdateParser(t *testing.T) {

	payload := "634|S|32"
	msg := Parse(payload)

	expected := Message{Payload: payload, Sequence: 634, Type: "S", From: "32"}

	if msg != expected {
		t.Fail()
	}
}
