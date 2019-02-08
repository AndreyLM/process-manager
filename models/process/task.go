package process

import "time"

// Process - sub process about process
type Process struct {
	Data   *map[int]interface{}
	Status string
	Expire time.Time
}
