package task

import (
	"fmt"
)

// Collection - collection of tasks
type Collection struct {
	Tasks []*Task
}

// CreateCollection - creates new collection
func CreateCollection() *Collection {
	return &Collection{}
}

// TaskExist - checks if task exist
func (c *Collection) TaskExist(name string) bool {
	for _, val := range c.Tasks {
		if val.Name == name {
			return true
		}
	}

	return false
}

// AddTask - add task to collection
func (c *Collection) AddTask(t *Task) error {
	for _, val := range c.Tasks {
		if val.Name == t.Name {
			return fmt.Errorf("Task \"%s\" already exist", t.Name)
		}
	}

	c.Tasks = append(c.Tasks, t)
	return nil
}

// RemoveTask - removes task from collection
func (c *Collection) RemoveTask(name string) {
	var newTasks []*Task
	for _, val := range c.Tasks {
		if val.Name != name {
			newTasks = append(newTasks, val)
		}
	}

	c.Tasks = newTasks
}

// GetTask - get task by name
func (c *Collection) GetTask(name string) *Task {
	for _, task := range c.Tasks {
		if task.Name == name {
			return task
		}
	}

	return nil
}

// RemoveExpired - removes expired processes from tasks
func (c *Collection) RemoveExpired() {
	for _, task := range c.Tasks {
		task.RemoveExpiredProcesses()
	}
}
