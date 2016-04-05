package main

import (
	"fmt"

	parser "github.com/lorchaos/gmaze/parser"
)

func main() {

	eventChannel := make(chan string)

	go parser.StartEventServer(eventChannel)

	state := parser.NewState()

	peerServer := parser.NewPeerServer()

	queue := parser.NewQueue()

	for txt := range eventChannel {

		msg := parser.Parse(txt)

		for _, m := range queue.Add(msg) {

			peerServer.Deliver(state.Process(m))
		}

		fmt.Printf("Received %d queue %d [%s]\n", msg.Sequence, queue.ExpectedSequence, msg.Payload)
	}
}
