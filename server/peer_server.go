package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

type PeerServer struct {
	peers map[string]Peer
}

type Peer struct {
	id      string
	conn    net.Conn
	channel chan string
}

func NewPeerServer() PeerServer {

	ln, err := net.Listen("tcp", ":9099")
	if err != nil {
		fmt.Println("Unable to start peer server [9099]", err)
		os.Exit(-1)
	}
	fmt.Println("Started peer server, port 9099")

	s := PeerServer{peers: make(map[string]Peer)}

	go s.process(ln)

	return s
}

func (p PeerServer) process(ln net.Listener) {

	for {
		socket, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting peers ", err)
		}

		fmt.Println("Accepting peers, count ", len(p.peers))

		if peer, err := NewPeer(socket); err == nil {
			p.peers[peer.id] = peer
		}
	}
}

func (p PeerServer) Deliver(msg string, predicate Predicate) {

	for k, v := range p.peers {
		if predicate(k) {
			v.Deliver(msg)
		}
	}
}

func NewPeer(conn net.Conn) (Peer, error) {

	fmt.Println("New peer connected")

	p := Peer{conn: conn, channel: make(chan string)}

	scanner := bufio.NewScanner(conn)

	if scanner.Scan() {
		p.id = scanner.Text()
	} else if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
		return Peer{}, err
	}

	fmt.Println("Registered peer ", p.id)

	go p.listen()

	return p, nil
}

func (p Peer) listen() {

	for payload := range p.channel {
		fmt.Printf("Delivering %s : %s\n", p.id, payload)
		io.WriteString(p.conn, payload)
		io.WriteString(p.conn, "\n")
	}
}

func (p Peer) Deliver(payload string) {
	p.channel <- payload
}
