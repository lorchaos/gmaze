package gmaze

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



func TestDelivery(t *testing.T) {
	
	one := Message { Payload: "", Sequence: 0, Type: "F", From: "60", To: "50" }

	d := func(m Message) { t.Logf("Got ", m) }

	queue := NewQueue(d)

	queue.Add(one)

	if queue.ExpectedSequence != 1 {
		t.Error("message was not delivered")
	}
}