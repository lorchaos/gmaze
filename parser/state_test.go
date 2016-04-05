package parser

import (
	"testing"
	"reflect"
	)


func TestBBr(t *testing.T) {
	
	msg := Message { Payload: "abcd", Sequence: 666, Type: "F", From: "60", To: "50" }

	state := NewState()

	d := state.Process(msg)
	
	if d.Payload != "abcd" {
		t.Fail()
	}

	if !reflect.DeepEqual(d.Target, []string {"50"}) {
		t.Fail()
	}
}