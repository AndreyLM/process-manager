package task

import (
	"errors"
	"fmt"
	"time"
)

// Task - model for task
type Task struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Processes   []*Process  `json:"processes"`
	Data        interface{} `json:"data"`
	MaxCount    int         `json:"maxCount"`
	Count       int         `json:"count"`
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
	t.RemoveExpiredProcesses()
	if t.Count < t.MaxCount {
		t.Count++
		t.Processes = append(t.Processes, p)
		return nil
	}

	return fmt.Errorf("Cannot add process to task. You reached max count: %d", t.MaxCount)
}

// DeleteProcess - deletes process from task
func (t *Task) DeleteProcess(processUUID string) (err error) {
	t.RemoveExpiredProcesses()
	var newProcesses []*Process

	for _, val := range t.Processes {
		if val.UUID.String() != processUUID {
			newProcesses = append(newProcesses, val)
		}
	}

	if len(newProcesses) == len(t.Processes) {
		err = errors.New("Process with such UUID does not exist or expired")
		return
	}

	t.Processes = newProcesses
	t.Count = len(t.Processes)
	return nil
}

// RemoveExpiredProcesses - removes expired processes
func (t *Task) RemoveExpiredProcesses() {
	var newProcesses []*Process

	for _, val := range t.Processes {
		if val.Expire.Time.UnixNano() > time.Now().UnixNano() {
			newProcesses = append(newProcesses, val)
		}
	}

	t.Processes = newProcesses
	t.Count = len(t.Processes)
}
