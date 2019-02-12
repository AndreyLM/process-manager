package task

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Process - sub process about process
type Process struct {
	UUID   uuid.UUID   `json:"uuid"`
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
	Expire CustomTime  `json:"expire"`
}

// CreateProcess - creates new process
func CreateProcess(duration int, data interface{}) (process *Process, err error) {
	var id uuid.UUID
	id, err = uuid.NewV4()
	if err != nil {
		return
	}
	expireAt := time.Now().Add(time.Duration(duration) * time.Second)
	process = &Process{
		UUID:   id,
		Data:   data,
		Status: "",
		Expire: CustomTime{Time: expireAt},
	}
	return
}
