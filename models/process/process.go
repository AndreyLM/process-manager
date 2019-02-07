package process

import (
	"fmt"
	"time"
)

// Process - model for process
type Process struct {
	Name        string
	Description string
	Tasks       []*Task
	MaxCount    int
	count       int
}

// CreateProcess - creates new process
func CreateProcess(name, description string, maxCount int) *Process {
	return &Process{
		Name:        name,
		Description: description,
		MaxCount:    maxCount,
	}
}

// AddTask - adds task to process
func (p *Process) AddTask(task *Task) error {
	p.removeExpiredTaskes()
	if p.count < p.MaxCount {
		p.count++
		p.Tasks = append(p.Tasks, task)
		return nil
	}

	return fmt.Errorf("Cannot add task process. You reached max count: %d", p.MaxCount)
}

func (p *Process) removeExpiredTaskes() {
	var newTasks []*Task

	for _, val := range p.Tasks {
		if val.Expire.UnixNano() > time.Now().UnixNano() {
			newTasks = append(newTasks, val)
		}
	}

	p.Tasks = newTasks
	p.count = len(p.Tasks)
}
