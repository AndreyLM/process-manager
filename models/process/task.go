package process

import (
	"fmt"
	"time"
)

// Task - model for task
type Task struct {
	Name        string
	Description string
	Processes   []*Process
	Data        *map[interface{}]interface{}
	MaxCount    int
	Count       int
}

// CreateTask - creates new task
func CreateTask(name, description string, maxCount int) *Task {
	return &Task{
		Name:        name,
		Description: description,
		MaxCount:    maxCount,
	}
}

// AddProcess - adds process to task
func (t *Task) AddProcess(p *Process) error {
	t.removeExpiredTaskes()
	if t.Count < t.MaxCount {
		t.Count++
		t.Processes = append(t.Processes, p)
		return nil
	}

	return fmt.Errorf("Cannot add process to task. You reached max count: %d", t.MaxCount)
}

func (t *Task) removeExpiredTaskes() {
	var newProcesses []*Process

	for _, val := range t.Processes {
		if val.Expire.UnixNano() > time.Now().UnixNano() {
			newProcesses = append(newProcesses, val)
		}
	}

	t.Processes = newProcesses
	t.Count = len(t.Processes)
}
