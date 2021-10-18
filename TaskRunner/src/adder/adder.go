package adder

import (
	"Profiling/TaskRunner/src/task_runner"
	"errors"
)

type AddTask struct {
	task_runner.Task
	x, y   int
	Result int
}

func (at *AddTask) Run() {

	at.Result = at.x + at.y
	at.SetDone(true)
}

func (at *AddTask) GetResult() (interface{}, error) {
	if at.GetDone() {
		return at.Result, nil
	} else {
		return 0, errors.New("Task is not done!")
	}
}

func Initialize(x, y int, name string) *AddTask {
	newTask := new(AddTask)
	newTask.Name = name
	newTask.SetDone(false)
	newTask.x = x
	newTask.y = y
	newTask.Result = 0
	newTask.SetDone(false)
	return newTask
}
