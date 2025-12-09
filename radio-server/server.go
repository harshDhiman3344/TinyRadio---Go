package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server on :8080")

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {

		fmt.Println("Error on server", err)

	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error in connection: ", err)
			continue
		}

		fmt.Println("New client joined: ", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Recieved: ", text)

		//Send the message back to client that u have been recieved

		conn.Write([]byte("You said: " + text + "\n"))

	}
}
