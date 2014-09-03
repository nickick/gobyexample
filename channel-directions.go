package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// this was confusing as hell for me - way I'm remembering this now: replace chan keyword with 'function body'
// then receiving and sending makes sense -
// Receiver: you have a receiver when the function body sends to it (pings in pong
// Sender: you have a sender when the function body receives from it (pings in ping, pongs in pong)
