
---
# Go Task Scheduler

This is a simple task scheduler implemented in Go. It uses Go's concurrency.

## Features

- Simple to use and extend
- Uses Go's goroutines for efficient concurrent task execution

## Getting Started

There is implementation of `Task` interface in your tasks. Here's an example of a task that prints a message every specified interval:

```go
type PrintTask struct {
    Message   string
    Duration  time.Duration
}

func (p *PrintTask) Execute() error {
    fmt.Println(p.Message)
    return nil
}

func (p *PrintTask) GetDuration() time.Duration {
    return p.Duration
}
```

Then, create a `TaskScheduler`, add your tasks to it and start the scheduler:

```go
func main() {
    taskScheduler := NewTaskScheduler()

    task1 := &PrintTask{
        Message:  "Runing fast!",
        Interval: 1 * time.Second,
    }

    task2 := &PrintTask{
        Message:  "Going slow",
        Interval: 2 * time.Second,
    }

    taskScheduler.AddTask(task1)
    taskScheduler.AddTask(task2)

    taskScheduler.Start()

    time.Sleep(15 * time.Second)
}
```

## How it works

Each task is run in its own goroutine, allowing for concurrent execution. The `TaskScheduler` uses the `time.Sleep` function to wait for the specified interval before executing each task.
