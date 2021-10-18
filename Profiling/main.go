package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func funcA() {
	fmt.Println("funcA")
	funcA1()
}

func funcA1() {
	fmt.Println("funcA1")
	funcA11()
}

func funcA11() {
	fmt.Println("funcA11")
}

func funcB() {
	fmt.Println("funcB")
	funcB1()
	funcB2()
}

func funcB1() {
	fmt.Println("funcB1")
}

func funcB2() {
	fmt.Println("funcB2")
}

func funcC() {
	fmt.Println("funcC")
}

// Run with
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	funcA()
	funcB()
	funcC()
}
