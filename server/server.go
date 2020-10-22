package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
	fmt.Println("Oopsie connection lost")
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	for {
		conn, _ := ln.Accept()
		conns <- conn
	}
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	reader := bufio.NewReader(client)
	for {
		// Read in new messages as delimited by '\n's
		msg, err := reader.ReadString('\n')
		if err != nil {
			handleError(err)
			break
		}
		// Tidy up each message and add it to the messages channel,
		mssg := Message{clientid, msg}
		// recording which client it came from.
		msgs <- mssg
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()
	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)
	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)
	//Start accepting connections
	go acceptConns(ln, conns)
	id := 1
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			clients[id] = conn
			// - start to asynchronously handle messages from this client
			go handleClient(clients[id], id, msgs)
			id++
		case msg := <-msgs:
			//TODO Deal with a new message
			for key, element := range clients {
				if key != msg.sender {
					fmt.Fprint(element, msg.message)
				}
			}
			// Send the message to all clients that aren't the sender
		}
	}
}
