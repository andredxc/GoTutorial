package main

import (
	"Profiling/TaskRunner/src/adder"
	"Profiling/TaskRunner/src/subtractor"

	"Profiling/TaskRunner/src/task_runner"
	"fmt"
)

func main() {

	at := adder.Initialize(10, 20, "Sum of 10 + 20")
	st := subtractor.Initialize(5, 11, "Subtraction of 5 - 11")
	tasks := []task_runner.TaskRunner{at, st}

	// TODO: TaskRunner metadata instead of redeclaring

	for _, runner := range tasks {
		// TODO: Run tasks as goroutines
		runner.Run()

		result, err := runner.GetResult()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		// fmt.Printf("Task=%s; done=%t; Result=%d\n", runner.Name, runner.IsDone(), result)
		fmt.Printf("Name=%s; done=%t; Result=%d\n", runner.GetName(), runner.GetDone(), result)
	}
}
