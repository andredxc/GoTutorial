package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Write an interactive console program that asks the user for the polar coordinates of a 2-dimensional point (radius and angle (degrees)).
Calculate the corresponding Cartesian coordinates x and y, and print out the result.
Use structs called polar and Cartesian to represent each coordinate system. Use channels and a goroutine:

A channel1 to receive the polar coordinates
A channel2 to receive the Cartesian coordinates
The conversion itself must be done with a goroutine, which reads from channel1 and sends it to channel2.
In reality, for such a simple calculation it is not worthwhile to use a goroutine and channels, but this solution would be quite appropriate for heavy computation.

*/

type polar struct {
	radius, angle float64
}

type cartesian struct {
	x, y float64
}

func main() {

	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {

	answers := make(chan cartesian)

	go func() {
		for {
			polarCoord := <-questions
			radAngle := polarCoord.angle * math.Pi / 180
			x := polarCoord.radius * math.Cos(radAngle)
			y := polarCoord.radius * math.Sin(radAngle)
			answers <- cartesian{x, y}
		}
	}()

	return answers
}

func interact(questions chan polar, answers chan cartesian) {

	var polarCoord polar
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter polar coordinates (radius, angle): ")

		line, err := reader.ReadString('\n')
		if err != nil {
			panic("Could not read string from terminal")
		}

		if numbers := strings.Fields(line); len(numbers) == 2 {

			radius, err := strconv.ParseFloat(numbers[0], 64)
			if err != nil {
				panic("Could not parse radius")
			}

			angle, err := strconv.ParseFloat(numbers[1], 64)
			if err != nil {
				panic("Could not parse angle")
			}
			polarCoord = polar{radius, angle}
		}

		fmt.Printf("Values read, radius=%f, angle=%f\n", polarCoord.radius, polarCoord.angle)
		questions <- polarCoord
		cartCoord := <-answers
		fmt.Printf("Equivalent cartesian coordinates: x=%f, y=%f\n", cartCoord.x, cartCoord.y)
	}
}
