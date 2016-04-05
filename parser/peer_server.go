package parser

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type PeerServer struct {
	peers map[string]Peer
}

type Peer struct {
	id   string
	conn net.Conn
}

func NewPeerServer() PeerServer {

	ln, err := net.Listen("tcp", ":9099")
	if err != nil {
		// handle error
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

		peer := NewPeer(socket)
		p.peers[peer.id] = peer
		fmt.Println("Ok")
	}
}

func (p PeerServer) Deliver(msg MessageDelivery) {

	if msg.Target == nil {

		for k, _ := range p.peers {
			p.deliver(k, msg.Payload)
		}

	} else {

		for k, _ := range msg.Target {
			p.deliver(k, msg.Payload)
		}
	}
}

func (p PeerServer) deliver(id string, payload string) {

	if peer, ok := p.peers[id]; ok {
		peer.Deliver(payload)
	}
}

func NewPeer(conn net.Conn) Peer {

	fmt.Println("New peer connected")

	p := Peer{conn: conn}

	scanner := bufio.NewScanner(conn)

	if scanner.Scan() {
		p.id = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}

	fmt.Println("Registered peer ", p.id)

	return p
}

func (p Peer) Deliver(payload string) {

	fmt.Printf("Delivering %s : %s\n", p.id, payload)

	io.WriteString(p.conn, payload)
	io.WriteString(p.conn, "\n")
}
