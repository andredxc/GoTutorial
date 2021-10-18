## Interfaces And Reflection

**Declaring interfaces and their implementation**
```
// Structs
type Square struct {
  side float32
}

type Circle struct {
  radius float32
}

// Interface
type Shaper interface {
  Area() float32
}

// Interface implementation
func (sq *Square) Area() float32 {
  return sq.side * sq.side
}
```
**Converting the type of an interface variable**
```
if v, ok := varI.(T); ok { // checked type assertion
    // If this conversion is valid, v will contain the value of varI converted to type T and ok will be true.
    Process(v)
    return
}
```

**Type Switch**
```
switch t := areaIntf.(type) {
    case *Square:
        fmt.Printf("Type Square %T with value %v\n", t, t)

    case *Circle:
        fmt.Printf("Type Circle %T with value %v\n", t, t)

    default:
        fmt.Printf("Unexpected type %T", t)
}
```

**Reading and Writing**
```
func main() {
  fmt.Println("Please enter your full name: ")
  fmt.Scanln(&firstName, &lastName)
  // fmt.Scanf("%s %s", &firstName, &lastName)
  fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
  fmt.Sscanf(input, format, &f, &i, &s)
  fmt.Println("From the string we read: ", f, i, s)
}
```

**Error Handling and Testing**
```
func Sqrt(f float64) (float64, error) {
  if f < 0 {
    return 0, errors.New("math - square root of negative number")
  }
  // implementation of Sqrt
}

if f, err := Sqrt(-1); err != nil {
  fmt.Printf("Error: %s\n", err)
}
```

**Parsing Flags Easily**
```
var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) { left <- 1+<-right }

func main() {
  flag.Parse()
  leftmost := make(chan int)
  var left, right chan int = nil, leftmost
  for i := 0; i < *ngoroutine; i++ {
    left, right = right, make(chan int)
    go f(left, right)
  }
  right <- 0 // start the chaining
  x := <-leftmost // wait for completion
  fmt.Println(x) // 100000, approx. 1.5 s
}
```

**Simple benchmarking of goroutine**
```
package main
import (
"fmt"
"testing"
)

func main() {
  fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
  fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(b *testing.B) { // makes buffered channel
  ch := make(chan int)
  go func() {
    for i := 0; i < b.N; i++ {
      ch <- i
    }
    close(ch)
  }()
  for _ = range ch { // iterating over channel without doing anything
  }
}

func BenchmarkChannelBuffered(b *testing.B) { // makes buffered channel with capacity of 128
  ch := make(chan int, 128)
  go func() {
    for i := 0; i < b.N; i++ {
      ch <- i
    }
    close(ch)

  }()
  for _ = range ch {
  }
}
```