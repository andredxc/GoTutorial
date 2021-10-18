package subtractor

import (
	"Profiling/TaskRunner/src/task_runner"
	"errors"
)

type SubtractTask struct {
	task_runner.Task
	x, y   int
	Result int
}

func (st *SubtractTask) Run() {

	st.Result = st.x - st.y
	st.SetDone(true)
}

func (st *SubtractTask) GetResult() (interface{}, error) {
	if st.GetDone() {
		return st.Result, nil
	} else {
		return 0, errors.New("Task is not done!")
	}
}

func Initialize(x, y int, name string) *SubtractTask {
	newTask := new(SubtractTask)
	newTask.Name = name
	newTask.x = x
	newTask.y = y
	newTask.Result = 0
	newTask.SetDone(false)
	return newTask
}
