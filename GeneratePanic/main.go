package main

import "log"

func badCall(zero int) int {
	return 12 / zero
}

func test() {
	defer func() {
		log.Println("Done")
		if err := recover(); err != nil {
			log.Printf("Run time panic %v\n", err)
		}
	}()
	badCall(0)
}

func main() {
	test()
}
