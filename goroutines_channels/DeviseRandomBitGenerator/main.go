package main

import "fmt"

const NUM_DIGITS = 100000

func RandomBinary() string {

	var result string

	zeroChannel := make(chan string)
	oneChannel := make(chan string)

	go makeString(zeroChannel, "0")
	go makeString(oneChannel, "1")

	for i := 0; i < NUM_DIGITS; i++ {
		select {
		case one := <-oneChannel:
			result += one
		case zero := <-zeroChannel:
			result += zero
		}
	}

	return result
}

func makeString(ch chan string, digit string) {
	for {
		ch <- digit
	}
}

func main() {
	result := RandomBinary()
	fmt.Printf("Done: %s\n", result)
}
