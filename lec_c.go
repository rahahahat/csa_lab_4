package main
import {
	"net"
	"fmt"
	"bufio"
	"os"
}

func read(*conn net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dail("tcp", "127.0.0.1:3080")
	for {
		fmt.Println("Enter text:")
		text, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, text)
		read(&conn)
	}
}