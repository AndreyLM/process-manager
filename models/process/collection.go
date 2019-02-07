package process

import "fmt"

// Collection - collection of processes
type Collection struct {
	processes []*Process
}

// CreateCollection - creates new collection
func CreateCollection() *Collection {
	return &Collection{}
}

// ProcessExist - checks if process exist
func (c *Collection) ProcessExist(name string) bool {
	for _, val := range c.processes {
		if val.Name == name {
			return true
		}
	}

	return false
}

// AddProcess - add process to collection
func (c *Collection) AddProcess(process *Process) error {
	for _, val := range c.processes {
		if val.Name == process.Name {
			return fmt.Errorf("Process %s already exist", process.Name)
		}
	}

	c.processes = append(c.processes, process)
	return nil
}

// RemoveProcess - removes process from collection
func (c *Collection) RemoveProcess(name string) {
	for i, val := range c.processes {
		if val.Name == name {
			c.processes = append(c.processes[:i], c.processes[i+1])
			return
		}
	}
}

// GetProcess - get process by name
func (c *Collection) GetProcess(name string) *Process {
	for _, process := range c.processes {
		if process.Name == name {
			return process
		}
	}

	return nil
}
