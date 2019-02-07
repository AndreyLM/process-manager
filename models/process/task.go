package process

import "time"

// Task - sub process about process
type Task struct {
	Data   *map[int]interface{}
	Status string
	Expire time.Time
}
