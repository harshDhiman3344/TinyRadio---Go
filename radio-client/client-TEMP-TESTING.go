package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server!")

	// Goroutine to read messages from server
	go func() {
		serverScanner := bufio.NewScanner(conn)
		for serverScanner.Scan() {
			fmt.Println("Server:", serverScanner.Text())
		}
	}()

	// Main loop to send messages to server
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("")
		if inputScanner.Scan() {
			clientMsg := inputScanner.Text()
			conn.Write([]byte(clientMsg + "\n"))
		}
	}
}
