package parser

import (
	"bufio"
	"fmt"
	"net"
)

type EventProcessor func(string)

func StartEventServer(evChannel chan string) {

	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		// handle error
	}

	fmt.Println("Started event server, port 9090")

	for {
		socket, err := ln.Accept()
		if err != nil {

			// handle error
		}

		fmt.Println("New event source connected")

		eventCount := processEventSource(socket, evChannel)

		fmt.Println("All events processed: %d", eventCount)
	}
}

func processEventSource(conn net.Conn, evChannel chan string) int {

	evCount := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		evChannel <- scanner.Text()
		evCount = evCount + 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}

	return evCount
}
