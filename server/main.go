package main

import (
	"fmt"
)

func main() {

	state := NewState()

	peerServer := NewPeerServer()

	queue := NewQueue()

	for txt := range StartEventServer() {

		msg := Parse(txt)

		for _, m := range queue.Add(msg) {

			predicate := state.Process(m)

			peerServer.Deliver(m.Payload, predicate)
		}

		fmt.Printf("Received %d queue %d [%s]\n", msg.Sequence, queue.ExpectedSequence, msg.Payload)
	}
}
