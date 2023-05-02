package main

import (
	"fmt"
	"time"
)

type Task interface {
	Execute() error
	GetDuration() time.Duration
}

// task implementation
// Struct that implements Task interface
type PrintTask struct {
	Message  string
	Duration time.Duration
}

// struct PrintTask implements Task interface by implementing Execute() and GetDuration() methods
func (p *PrintTask) Execute() error {
	fmt.Println(p.Message)
	return nil
}

func (p *PrintTask) GetDuration() time.Duration {
	return p.Duration
}

//end of task implementation

type TaskScheduler struct {
	tasks []Task
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{tasks: []Task{}}
}

func (s *TaskScheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

func (s *TaskScheduler) Start() {
	//iteration over all the tasks
	for _, task := range s.tasks {
		//start a goroutine for each task
		go func(t Task) {
			for {
				//in the loop, there is sleep for the duration of the task and afterward execute it
				time.Sleep(t.GetDuration())
				t.Execute()
			}
		}(task)
	}
}

func main() {
	scheduler := NewTaskScheduler()
	scheduler.AddTask(&PrintTask{
		Message:  "Running fast",
		Duration: 1 * time.Second,
	})
	scheduler.AddTask(&PrintTask{
		Message:  "Going slow",
		Duration: 4 * time.Second,
	})
	scheduler.Start()

	time.Sleep(15 * time.Second)
}
