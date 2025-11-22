package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Tiny chat starting...")
	fmt.Println("Hello this is a test again for this tiny chat project, i am gonna learn Go lang using this project.")

	// Make channels and scanner that carries the strings.
	messages := make(chan string) // Receiver 1 channel
	MainPipe := make(chan string) // Sender → Processor
	rec2 := make(chan string)     // Receiver 2 channel
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You: ")

	// MAKING SLICE AND RECURSING OVER THE RECIEVERS APPENDED TO IT:
	recievers := []chan string{}

	recievers = append(recievers, messages)
	recievers = append(recievers, rec2)

	// Sender Goroutine
	go func() {
		for {
			if scanner.Scan() {
				text := scanner.Text()
				MainPipe <- text
			}
		}
	}()

	// PROCESSOR GOROUTINE.
	go func() {
		for {
			mainMsg := <-MainPipe

			//Check Empty msg
			if mainMsg == "" {
				continue
			}

			// broadcast to all receivers
			for _, r := range recievers {
				r <- mainMsg
			}
		}
	}()

	// reciever2
	go func() {
		for {
			msg2 := <-rec2
			fmt.Printf("[%s]Reciever 2: %s\n", time.Now().Format("15:04:05"), msg2)
			fmt.Println("You: ")
		}
	}()

	// reciever1
	for {
		msg := <-messages
		fmt.Printf("\n[%s]Reciever 1: %s\n", time.Now().Format("15:04:05"), msg)
		fmt.Println("You: ")
	}
}
