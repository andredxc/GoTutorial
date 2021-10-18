package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

var nrchars, nrwords, nrlines int

func main() {

    inputReader := bufio.NewReader(os.Stdin)

    bDone := false

    for !bDone {
        fmt.Println("Input:")
        input, err := inputReader.ReadString('\n')
        if (input == "S") {
            bDone = true

        } else if err == nil{
            fmt.Printf("Input: %s\n", input)
            Counters(strings.TrimSuffix(input, "\n"))
        }
    }

    fmt.Printf("Done, nrchars=%d, nrwords=%d, nrlines=%d\n", nrchars, nrwords, nrlines)
}

func Counters(input string) {

    nrchars += len(strings.Replace(input, " ", "", -1))
    nrwords += len(strings.Split(input, " "))
    nrlines += 1

    fmt.Printf("Counters: line=%s nrchars=%d, nrwords=%d, nrlines=%d\n", strings.TrimSpace(input), nrchars, nrwords, nrlines)
}