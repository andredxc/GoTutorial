package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

/*
	Write a program that starts a goroutine sum to perform a sum of 2 integers and then waits on the result to print it.
*/

func sum(x, y int, c chan int) {
	// time.Sleep(1e9)
	c <- x + y
}

func main() {

	if len(os.Args) < 3 {
		panic(errors.New("Not enough arguments"))
	}

	var x, y int
	var err error

	if x, err = strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("Error reading int from %s\n", os.Args[1])
	}

	if y, err = strconv.Atoi(os.Args[2]); err != nil {
		fmt.Printf("Error reading int from %s\n", os.Args[2])
	}

	ch := make(chan int, 1)
	go sum(x, y, ch)
	fmt.Println("Waiting for result")
	fmt.Printf("%d + %d = %d\n", x, y, <-ch)
}
