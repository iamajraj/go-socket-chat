package main

import (
	"fmt"
	"net"
)

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = "8000"
)

var connections []net.Conn

func main() {

	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

	if err != nil {
		fmt.Println("Something went wrong while listening the server")
	}

	fmt.Println("Server running on port 8000")

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Connection failed")
		}

		fmt.Println("Got a client's connection")
		connections = append(connections, conn)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var result = make([]byte, 1024)

	for {
		mlen, err := conn.Read(result)

		if err != nil {
			fmt.Println("Failed to read")
			break
		}
		msg := string(result[:mlen])
		fmt.Println(msg)
		go sendMessage(msg)
	}

}

func sendMessage(msg string) {
	for _, conn := range connections {
		conn.Write([]byte(msg))
	}
}
