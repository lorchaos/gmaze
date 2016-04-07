package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartEventServer() chan string {

	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Unable to open socket", err)
		os.Exit(-10)
	}

	fmt.Println("Started event server, port 9090")

	eventChannel := make(chan string)

	go listen(ln, eventChannel)

	return eventChannel
}

func listen(ln net.Listener, evChannel chan string) {

	socket, err := ln.Accept()
	if err != nil {
		fmt.Println("Unable to accept socket", err)
		os.Exit(-11)
	}

	defer socket.Close()

	fmt.Println("New event source connected")

	processEventSource(socket, evChannel)

	close(evChannel)

	fmt.Println("All events processed")
}

func processEventSource(conn net.Conn, evChannel chan string) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		evChannel <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
}
