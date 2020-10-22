package main

import {
	"bufio"
	"net"
	"fmt"
}
func handleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	for {
		msg,_ := reader.ReadString('\n')
		fmt.Printf(msg)
		fmt.Fprintln(*conn, "OK")
	}
}

func main() {
	ls,_ := net.Listen("tcp", ":8030")
	for {
		conn,_ := ls.Accept()
		go handleConnection(*conn)
	}
}