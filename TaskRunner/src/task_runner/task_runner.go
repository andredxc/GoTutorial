package task_runner

type TaskRunner interface {
	Run()
	GetDone() bool
	SetDone(v bool)
	GetResult() (interface{}, error)
	GetName() string
}

type Task struct {
	Name string
	Done bool
}

func (t *Task) GetName() string {
	return t.Name
}

func (t *Task) GetDone() bool {
	return t.Done
}

func (t *Task) SetDone(val bool) {
	t.Done = val
}
