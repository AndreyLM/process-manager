package task

import (
	"encoding/json"
	"fmt"
	"time"
)

// CustomTime - time for process
type CustomTime struct {
	Time time.Time `json:"time"`
}

// MarshalJSON - custom marshaling time
func (t CustomTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("%s", t.Time.Format("02/01/2006 15:04:05 .9999"))
	return json.Marshal(stamp)
}
