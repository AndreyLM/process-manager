package task

import "time"

// Process - sub process about process
type Process struct {
	Data   *map[interface{}]interface{}
	Status string
	Expire time.Time
}
