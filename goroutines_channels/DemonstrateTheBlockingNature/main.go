package main

import (
	"log"
	"time"
)

func main() {

	ch := make(chan int)

	go send(ch)
	go receive(ch)
	time.Sleep(6 * 1e9)
	log.Println("Done!")
}

func send(ch chan int) {
	log.Println("Sending integer though channel")
	ch <- 1
	log.Println("Integer sent")
}

func receive(ch chan int) {

	log.Println("Waiting to receive")
	time.Sleep(5 * 1e9)
	log.Printf("Received %d\n", <-ch)
}
